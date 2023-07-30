# Full stack demo-intro

This is a modern full-stack demo developed independently by Leon LiangNing individual .

In terms of functionality, half of it is a feed and half is e-commerce, referenced from the XiaoHongShu[小红书 – 你的生活指南 on the App Store (apple.com)](https://apps.apple.com/us/app/%E5%B0%8F%E7%BA%A2%E4%B9%A6-%E4%BD%A0%E7%9A%84%E7%94%9F%E6%B4%BB%E6%8C%87%E5%8D%97/id741292507).

## projects

- [Golang-framework kit-demo](https://github.com/lightning560/mio) is named mio
- [Golang Backend microservice-demo](https://github.com/lightning560/go-microservice)
- [SRE & Full chain stress test](https://github.com/lightning560/mio/blob/main/SRE&Test.md)
- [Web-demo](https://github.com/lightning560/Web-demo)
- [iOS-demo](https://github.com/lightning560/iOS-demo)
- [Flutter-demo](https://github.com/lightning560/Flutter-demo)

# write and read heavy solution-syllabus

How to solve write heavy and read heavy problems in infra and Mio Golang framework.

For example, in a business system, there are write-heavy operations such as the history service and read-heavy operations such as the Sales Promotion service.

## infra and Golang framework solution

### cache

- Cache Aside data consistency
- fanout design pattern. improve performance
- redis big key problem.multicopy or multiclustering.
- cache penetration issues.using singleflight
- hot key
  - topk algorithm,use HeavyKeeper
  - The local cache with zero GC
  - redis 6.0 client side tracing to solve the data consistency problem of localcache
  - Using Interceptor and Command Design Patterns in a redis client in a Mio framework kit.

### db

- Google SRE Circuit breaking algorithm
- Distributed rate limiting,refer to Mio Quota section
- Null filtering, using redis-bloom module bloomfliter implementation

## Business system solution

### write heavy

- Write to message queue, using kafka or redis pub/sub
- write back, write to Redis first
- Batch writing, then flush to db

## read heavy

- Services separate isolation
- Staggered requests, so that users request in batches. Server stagger and App client stagger.
- Weaken peaks, increase request intervals
- Rate limiting, refer to the Mio golang framework.
- Cache, refer to the following cache section.

# cache

## refer

### doordash

<https://doordash.engineering/2018/08/03/avoiding-cache-stampede-at-doordash/>

### paper《scaling memcache at Facebook》

[usenix.org/system/files/conference/nsdi13/nsdi13-final170_update.pdf](https://www.usenix.org/system/files/conference/nsdi13/nsdi13-final170_update.pdf)

## Redis usage

Protobuf is used for serialization to minimize the size.

## Cache Aside Data Consistency

In the Cache Aside model, cache refilling operations when reading cache miss, synchronous cache updates when modifying data, and asynchronous cache compensation by message queues cannot satisfy "Happens Before", and there will be cases of overwriting each other.

Add operation, setnx operations do not overwrite.
Set operation, setex operations overwrite.

Simultaneous read/write operations:

1. Read operation, read cache, cache MISS.
2. Read operation, read DB, data read.
3. Write operation, update DB data.
4. Write operation SET Cache (can be asynchronous Job operation, Redis can use SETEX operation)
5. Read operation, ADD operation writes back cache (can be Job asynchronous operation, Redis can use SETNX operation)
The write operation uses the SET command to overwrite the write cache; the read operation uses the ADD operation to write back MISS data, thus ensuring that the latest data from the write operation is not overwritten by the write-back data from the read operation.

![](/img/Pasted%20image%2020230612095507.png)

## fanout design pattern

![](/img/Pasted%20image%2020230611221027.png)
Goroutine and Channel are used for adding data to the cache.
The primary purpose is to improve CPU utilization.

`example`

```go
func Example() {
 // 这里只是举个例子 真正使用的时候 rpc 传过来的context
 var c = context.Background()
 // 新建一个fanout 对象 名称为cache 名称主要用来上报监控和打日志使用 最好不要重复
 // (可选参数) worker数量为1 表示后台只有1个线程在工作
 // (可选参数) buffer 为1024 表示缓存chan长度为1024 如果chan满了 再调用Do方法就会报错 设定长度主要为了防止OOM
 cache := New("cache", Worker(1), Buffer(1024))
 // 需要异步执行的方法
 // 这里传进来的c里面的meta信息会被复制 超时会忽略 addCache拿到的context已经没有超时信息了
 cache.Do(c, func(c context.Context) { addCache(c, 0, 0) })
 // 程序结束的时候关闭fanout 会等待后台线程完成后返回
 cache.Close()
}
```

fanout code

```go
package fanout

import (
 "context"
 "errors"
 "runtime"
 "sync"
)
var (
 // ErrFull chan full.
 ErrFull = errors.New("fanout: chan full")
)

type options struct {
 worker int
 buffer int
}

// Option fanout option
type Option func(*options)

// Worker specifies the worker of fanout
func Worker(n int) Option {
 if n <= 0 {
  panic("fanout: worker should > 0")
 }
 return func(o *options) {
  o.worker = n
 }
}

// Buffer specifies the buffer of fanout
func Buffer(n int) Option {
 if n <= 0 {
  panic("fanout: buffer should > 0")
 }
 return func(o *options) {
  o.buffer = n
 }
}

type item struct {
 f   func(c context.Context)
 ctx context.Context
}

// Fanout async consume data from chan.
type Fanout struct {
 name    string
 ch      chan item
 options *options
 waiter  sync.WaitGroup

 ctx    context.Context
 cancel func()
}

// New new a fanout struct.
func New(name string, opts ...Option) *Fanout {
 if name == "" {
  name = "fanout"
 }
 o := &options{
  worker: 1,
  buffer: 1024,
 }
 for _, op := range opts {
  op(o)
 }
 c := &Fanout{
  ch:      make(chan item, o.buffer),
  name:    name,
  options: o,
 }
 c.ctx, c.cancel = context.WithCancel(context.Background())
 c.waiter.Add(o.worker)
 for i := 0; i < o.worker; i++ {
  go c.proc()
 }
 return c
}

func (c *Fanout) proc() {
 defer c.waiter.Done()
 for {
  select {
  case t := <-c.ch:
   wrapFunc(t.f)(t.ctx)
   // stats.State(c.name+"_channel", int64(len(c.ch)))
  case <-c.ctx.Done():
   return
  }
 }
}

func wrapFunc(f func(c context.Context)) (res func(context.Context)) {
 res = func(ctx context.Context) {
  defer func() {
   if r := recover(); r != nil {
    buf := make([]byte, 64*1024)
    buf = buf[:runtime.Stack(buf, false)]
   }
  }()
  f(ctx)
 }
 return
}

// Do save a callback func.
func (c *Fanout) Do(ctx context.Context, f func(ctx context.Context)) (err error) {
 if f == nil || c.ctx.Err() != nil {
  return c.ctx.Err()
 }
 select {
 case c.ch <- item{f: f, ctx: ctx}:
 default:
  err = ErrFull
 }
 return
}

// Close close fanout
func (c *Fanout) Close() error {
 if err := c.ctx.Err(); err != nil {
  return err
 }
 c.cancel()
 c.waiter.Wait()
 return nil
}
```

## Cache penetration

- Use distributed lock
Set a lock key, and only one person succeeds and returns, and this person is assigned to execute the source return operation. Other candidates poll the cache lock key. If it doesn't exist, they read the data cache. If the cache hits, they return. If the cache misses, they continue to grab the lock.
This is not recommended as distributed lock is troublesome.
- Use queue
If cache misses, hand it over to the queue to aggregate one key, load the data, and write back to the cache. For the current miss request, use singleflight to ensure sourcing, such as the comment architecture implementation. This is suitable for tasks with heavy data loading during sourcing, such as comment misses that only return the first page, but require the complete construction of comment data indexes.
- Use lease mechanism
Refer to the solution used by Facebook
By adding a lease mechanism, these two problems can be effectively avoided. The lease is a 64-bit token that is bound to the key of the client request. For outdated setting, the lease is verified during writing to solve this problem. For thundering herd, each key is allocated once every 10 seconds. If the client has not obtained a lease, it can wait a bit before accessing the cache. At this time, there is often already data in the cache. (Requires support from basic library & modification of cache source code)

### Use singleflight to prevent penetration

Consistently hash keywords so that keys of a certain dimension must hit a certain node, then use mutual exclusion locks within the node to ensure merger back to source. However, there is no solution for batch queries.

The main scene is similar to listByXID.

```go
var cacheSingleFlights = [2]*singleflight.Group{{}, {}}
```

Initialize as many as the number of calls.

```go
var ( 
 sfKey1 = "key1" 
 wg *sync.WaitGroup 
 sf singleflight.Group 
 nums = 10 
)
// getValueBySingleflight 使用singleflight取cacheKey对应的value值
func getValueBySingleflight(idx int, cacheKey string) (string, error) {
   log.Printf("idx %v into-cache...", idx)
   // 调用singleflight的Do()方法
   value, _, _ := sf.Do(cacheKey, func() (ret interface{}, err error) {
      log.Printf("idx %v is-setting-cache", idx)
      // 休眠0.1s以捕获并发的相同请求
      time.Sleep(100 * time.Millisecond)
      log.Printf("idx %v set-cache-success!", idx)
      return "myValue", nil
   })
   return value.(string), nil
}
```

# cache bigkey problem:redis

If the list is too long, it could also be a BigKey.

## Multiple key replicas

Data synchronization could be cumbersome as it's not so easy.
Each business does not necessarily have to use a single cluster. If memory is insufficient, each key can correspond to a separate cluster.
Hotspot issues still exist even with a single key. If a single replica isn't enough, like ms_1, ms_2, ms_3, a key can be dispersed into multiple clusters to avoid hot spots. But data consistency issues could be tricky to handle.

## Multiple Clusters

One key for one cluster, it's not necessary to use the same cluster for all.

## Analysis

### BigKeys command analysis

Redis 4.0 and later versions provide the --BigKeys command which can analyze the Top 1 BigKey for each type of data structure in the instance, while also providing the number of key-values for each data type and their average size.

### RDB file analysis

Open-source tools such as RDB-tools can be used to analyze the RDB files of a Redis instance and find out the BigKeys in there.

# cache hot Key problem

## Solutions

- Broadcast small tables and upgrade RemoteCache to LocalCache. The app is updated regularly, and even the operations platform can support the broadcast refresh of LocalCache.
- Actively monitor for preheating in defense.
- The Mio golang framework supports hotkey discovery and automatically short-lived localcache.
- Multi-cluster support;
- Multi-key design: use multiple replicas to reduce node hotkey issues.
  - Use multiple replicas such as ms_1, ms_2, ms_3. Each node stores a set of data, which helps to distribute requests across multiple nodes, avoiding single point hotspot issues.

## Proactive Preheating + Whitelist Mechanism

In the scenario of a sale Promotion, through near real-time traffic statistics, the topk of the most visited products are obtained, which are then proactively pushed to the application instances. After receiving the push, the application instance caches the hotspot room locally for access. Through near real-time statistics and push, most of the popular products can be intercepted.

In popular events, hotspots can be predicted in advance. The id related to the event will immediately become hot data as soon as the event starts. In this scenario, we can add the id related to the event to the whitelist in advance. When the event starts, as long as the id is accessed, it will be immediately added to the local cache, without waiting for the hotspot detection module to detect it. Whitelisting can avoid a sudden burst of hotspot traffic that could crash the backend storage. Although we have increased the detection sensitivity of hotspots by adding decay factors, whitelisting can further reduce the sudden impact of super hotspots on the backend.

```go
if ok := topk.Add(key, 1); ok {
    lru.Add(key, value, ttl)
    return
}
if ok := inWhileList(key); ok {
    lru.Add(key, value, ttl)
    return
} 
```

## Proactive Hotspot Monitoring

- topk
- min-heap
- redis client slide tracing
- local cache

## HeavyKeeper for Top-K

The HeavyKeeper algorithm can fetch topks accurately and memory-saving
Redis topk uses this HeavyKeeper algorithm.

HeavyKeeper is used to obtain highly accurate top-k calculation results with minimal memory overhead in streaming data. The top-k calculated represents the hot keys we want to know. HeavyKeeper was proposed in 2018 and has a similar overall structure to Count-Min Sketch. However, unlike Count-Min Sketch, each position in the two-dimensional array of HeavyKeeper does not only store the count value but also simultaneously stores the hash fingerprint and the count value. There are also differences in how they handle hash collisions. Count-Min Sketch directly increments the count value for hash collisions and then takes the minimum value of the results, which can cause the final statistics to be overly large. In contrast, HeavyKeeper applies probability decay to the count values in case of hash collisions.

The implementation of top-k is similar to Count-Min Sketch, using a min-heap to maintain the top-k elements.

### Paper

[https://www.usenix.org/system/files/conference/atc18/atc18-gong.pdf](https://www.usenix.org/system/files/conference/atc18/atc18-gong.pdf)

### Comparison

By comparing and analyzing the top-k accuracy, frequency accuracy, and memory overhead of these three algorithms, it is found that HeavyKeeper can achieve higher statistical accuracy with the same memory overhead. Therefore, we have decided to implement our hotspot monitoring based on HeavyKeeper.
![](/img/Pasted%20image%2020230613145046.png)

### implement

#### example

```go
topk := NewHeavyKeeper(10, 10000, 5, 0.925)
topk.Add(key, 1)

fmt.Println(topk.List()[0:5])
```

#### interface

```go
// Topk algorithm interface.
type Topk interface {
    // Add item and return if item is in the topk.
    Add(item string, incr uint32) bool
    // List all topk items.
    List() []Item
    // Expelled watch at the expelled items.
    Expelled() <-chan Item
}
```

`Add method`,processes each element in the stream and determines whether it can be added to the topk according to the algorithm.

- `List method`,returns a list of the elements currently in the topk.
- `Expelled method`,returns the element that is currently expelled from the topk. The business can customize the logic design based on the topk, such as the hotspot local cache mentioned below, when the data does not belong to the topk, this method can be used to expel the old hotspot data out of the local cache to ensure that the hotspot freshness of the local cache.

`add method`

```go
func (topk *HeavyKeeper) Add(key string, incr uint32) bool {
 keyBytes := []byte(key)
 itemFingerprint := murmur3.Sum32(keyBytes)
 var maxCount uint32

 // compute d hashes
 for i, row := range topk.buckets {

  bucketNumber := murmur3.Sum32WithSeed(keyBytes, uint32(i)) % uint32(topk.width)
  fingerprint := row[bucketNumber].fingerprint
  count := row[bucketNumber].count

  if count == 0 {
   row[bucketNumber].fingerprint = itemFingerprint
   row[bucketNumber].count = incr
   maxCount = max(maxCount, incr)

  } else if fingerprint == itemFingerprint {
   row[bucketNumber].count += incr
   maxCount = max(maxCount, row[bucketNumber].count)

  } else {
   for localIncr := incr; localIncr > 0; localIncr-- {
    var decay float64
    curCount := row[bucketNumber].count
    if row[bucketNumber].count < LOOKUP_TABLE {
     decay = topk.lookupTable[curCount]
    } else {
     // decr pow caculate cost
     decay = topk.lookupTable[LOOKUP_TABLE-1]
    }
    if topk.r.Float64() < decay {
     row[bucketNumber].count--
     if row[bucketNumber].count == 0 {
      row[bucketNumber].fingerprint = itemFingerprint
      row[bucketNumber].count = localIncr
      maxCount = max(maxCount, localIncr)
      break
     }
    }
   }
  }
 }
 minHeap := topk.minHeap.Min()
 if len(topk.minHeap.Nodes) == int(topk.k) && maxCount < minHeap {
  return false
 }
 // update minheap
 itemHeapIdx, itemHeapExist := topk.minHeap.Find(key)
 if itemHeapExist {
  topk.minHeap.Fix(itemHeapIdx, maxCount)
  return true
 }
 expelled := topk.minHeap.Add(&minheap.Node{Key: key, Count: maxCount})
 if expelled != nil {
  topk.expell(Item{Key: expelled.Key, Count: expelled.Count})
 }
 return true
}
func (topk *HeavyKeeper) List() []Item {
 items := topk.minHeap.Sorted()
 res := make([]Item, 0, len(items))
 for _, item := range items {
  res = append(res, Item{Key: item.Key, Count: item.Count})
 }
 return res
}

func (topk *HeavyKeeper) expell(item Item) {
 select {
 case topk.expelled <- item:
 default:
 }
}
```

## LocalCache

### Based on freecache

With expire time and zero-GC based on freecache as local cache
Second packaging, used for local caching and support metric reporting

Precautions:

- **The size of the cache key needs to be less than 65535, otherwise it cannot be stored in the local cache** (The key is larger than 65535)
- **The size of the cache value needs to be less than 1/1024 of the total cache capacity, otherwise it cannot be stored in the local cache** (The entry size need less than 1/1024 of cache size)

### LRU as the cache elimination algorithm for local cache

Topk writing into LRU, LFU itself has the attribute of detecting hotspots, If hot spot detection is used, LRU can be used

## local cache data consistency: Redis 6.0's client side tracing

Use Redis client tracing to get notifications of data changes in order to maintain data consistency.

Redis value's expiration notification

```go
// ClientTracking is a helper function that enables Redis client-side caching
func (hk *HotKey) NewClientTracking() {
 // Create connetion providing key invalidation callback.
 dialer := new(client.Dialer)
 // 失效通知回调
 dialer.InvalidateCallback = func(keys []string) {
  for _, key := range keys {
   hk.Remove(key)
   fmt.Printf("clear localCache %s\n", key)
  }
 }
 conn, err := dialer.Dial(hk.config.Address)
 if err != nil {
  log.Fatal(err)
 }

 broadcast := true
 if err := conn.ClientTracking(true, nil, hk.config.Prefix, broadcast, false, false, false).Err(); err != nil {
  log.Fatal(err)
 }
}
```

## Integration into Mio Golang Framework

### Interceptor + Command Design Pattern

```go
func filterInterceptor(compName string, config *config, logger *xlog.Logger) *interceptor {
 return newInterceptor(compName, config, logger).setAfterProcess(
  func(ctx context.Context, cmd redis.Cmder) error {
   fmt.Println("filterInterceptor cmd.Name():", cmd.Name())
   return nil
 })
}
```

# db

## Distributed rate limit Quota

refer to Mio Golang framework

## Circuit Breaker

refer to Mio Golang framework

## Null Value Filtering

Write cache policy after read failure (after degradation, generally read failure does not trigger rewrite cache)
Empty cache settings. For some data, the database may always be empty. In this case, an empty cache should be set to avoid every request cache miss hitting the DB directly.

### Algorithm Selection

Bloom filters can determine that a key definitely does not exist, but cannot judge that a key definitely exists, but can control the false positive rate to determine existence.
If it's a frequently deleting then use Cuckoo.
Cuckoo is faster, deletable, but takes up a bit more memory.
Finally choose bloomfilter.

### Use Redis-Bloom module

use github.com/RedisBloom/redisbloom-go library as redis client

```go
func (r *Redbloom) BfAdd(key string, item string) (exists bool, err error) {
 exists, err = r.Client.Add(key, item)
 if err != nil {
  log.Errorf("eredbloom add error", log.FieldErr(err), log.FieldKey(key), log.FieldValue(item))
 }
 return
}

func (r *Redbloom) BfExists(key string, item string) (exists bool, err error) {
 exists, err = r.Client.Exists(key, item)
 if err != nil {
  log.Errorf("eredbloom exists error", log.FieldErr(err), log.FieldKey(key), log.FieldValue(item))
 }
 return
}

// BfAddMulti -Adds one or more items to the Bloom Filter, creating the filter if it does not yet exist. args: key - the name of the filter item - One or more items to add
func (r *Redbloom) BfAddMulti(key string, items []string) (res []int64, err error) {
 res, err = r.Client.BfAddMulti(key, items)
 if err != nil {
  log.Errorf("eredbloom BfAddMulti error", log.FieldErr(err), log.FieldKey(key))
 }
 return
}

// BfExistsMulti - Determines if one or more items may exist in the filter or not. args: key - the name of the filter item - one or more items to check
func (r *Redbloom) BfExistsMulti(key string, items []string) (res []int64, err error) {
 res, err = r.Client.BfExistsMulti(key, items)
 if err != nil {
  log.Errorf("eredbloom BfExistsMulti error", log.FieldErr(err), log.FieldKey(key))
 }
 return
}

// BfInsert - This command will add one or more items to the bloom filter, by default creating it if it does not yet exist.
func (r *Redbloom) BfInsert(key string, cap int64, errorRatio float64, expansion int64, noCreate bool, nonScaling bool, items []string) (res []int64, err error) {
 res, err = r.Client.BfInsert(key, cap, errorRatio, expansion, noCreate, nonScaling, items)
 if err != nil {
  log.Errorf("eredbloom insert error", log.FieldErr(err), log.FieldKey(key))
 }
 return
}

func (r *Redbloom) CfAdd(key string, item string) (exists bool, err error) {
 exists, err = r.Client.CfAdd(key, item)
 if err != nil {
  log.Errorf("eredbloom add error", log.FieldErr(err), log.FieldKey(key), log.FieldValue(item))
 }
 return
}

func (r *Redbloom) CfAddMulti(key string, items []string) (res []int64, err error) {
 for _, item := range items {
  exists, err := r.Client.CfAdd(key, item)
  if err != nil {
   log.Warnf("eredbloom add error", log.FieldErr(err), log.FieldKey(key))
  }
  if exists {
   res = append(res, 1)
  } else {
   res = append(res, 0)
  }
 }
 return
}

func (r *Redbloom) CfExists(key string, item string) (exists bool, err error) {
 exists, err = r.Client.CfExists(key, item)
 if err != nil {
  log.Errorf("eredbloom CfExists error", log.FieldErr(err), log.FieldKey(key), log.FieldValue(item))
 }
 return
}

func (r *Redbloom) CfExistsMulti(key string, items []string) (res []int64, err error) {
 for _, item := range items {
  exists, err := r.Client.CfExists(key, item)
  if err != nil {
   log.Errorf("eredbloom CfExistsMulti error", log.FieldErr(err), log.FieldKey(key))
  }
  if exists {
   res = append(res, 1)
  } else {
   res = append(res, 0)
  }
 }
 return
}

// Adds one or more items to a cuckoo filter, allowing the filter to be created with a custom capacity if it does not yet exist.
func (r *Redbloom) CfInsert(key string, cap int64, noCreate bool, items []string) ([]int64, error) {
 res, err := r.Client.CfInsert(key, cap, noCreate, items)
 if err != nil {
  log.Errorf("eredbloom insert error", log.FieldErr(err), log.FieldKey(key))
 }
 return res, err
}
```

## CDC (Change Data Capture)

To ensure timeliness and consistency.

# History service: Write Heavy

For a write-heavy workload, the following strategies can be employed:

1. Last-Write-Wins (==last-write win==): Only the latest data for each ID needs to be persisted. In high-frequency user synchronization logic, it is sufficient to persist only the last write data.
2. Kafka + Consistent Hashing: Use Kafka as a message queue combined with consistent hashing to distribute the write load across multiple nodes, ensuring scalability and fault tolerance.
3. Write Back: Write the data to Redis and periodically flush a batch of data to the database. This approach combines in-memory storage with database persistence.
4. When reading data, if the data is not found in Redis, it is not necessary to retrieve it again and fill it back into Redis, as the old data is not needed.

It is important to consider the write capacity of the database. Sharding the data into 100-500 partitions using techniques like sharding and partitioning can help improve write performance.
If the workload becomes overwhelming for a single database, it may be necessary to consider using a distributed database solution to handle the increased load.

## Writing to Message Queue (MQ)

## Memory Write: Write Back

It is difficult to avoid data inconsistency issues. The approach is to first write the data to a history storage and then, after a certain number of operations or a certain period of time, write it to the database.

## Aggregated Write: Merge Batch Write

### Scenarios

- Count,Like count,View count
- Calculation of small gifts or rewards

In these scenarios, the values in Redis are aggregated and then flushed to the database using FlushCacheToDB. This approach addresses the issue of write hotspots. It is usually caused by a large number of users performing write operations on the same anchor, such as sending gifts or posting comments, which results in a high level of concurrent writes for a single record. Upon further analysis, we find that these concurrent write scenarios typically involve incremental or decremental operations on a single record, such as experience points, scores, likes, etc. These types of scenarios naturally support aggregation. Therefore, we have developed an SDK for aggregated writes that can aggregate the business's data change operations within a set period, such as combining three operations of +1, +2, and -1 into a single operation of +2. Implementing this SDK requires considering parameters such as the aggregation window size, downstream database pressure, and ensuring data consistency in case of service exceptions or restarts.

![](/img/Pasted%20image%2020230613140611.png)

### example

Continuously retry writing and flushing to the database. Initialize the merge process and write to the database based on time or data volume. Perform nightly reconciliation using cron. Note that not all business processes can be reconciled.

```go
package service

import (
 "context"
 "encoding/json"
 "fmt"
 "hash/crc32"
 "sort"
 "strings"
 "time"
)

func (s *Service) serviceConsumeproc() {
 var (
  err  error
  msgs = s.serviceHisSub.Messages()
 )
 for {
  msg, ok := <-msgs
  if !ok {
   log.Error("s.serviceConsumeproc closed")
   return
  }
  if s.c.Job.IgnoreMsg {
   err = msg.Commit()
   log.Info("serviceConsumeproc key:%s partition:%d offset:%d err: %+v, ts:%v ignore", msg.Key, msg.Partition, msg.Offset, err, msg.Timestamp)
   continue
  }
  ms := make([]*model.Merge, 0, 32)
  if err = json.Unmarshal(msg.Value, &ms); err != nil {
   log.Error("json.Unmarshal(%s) error(%v)", msg.Value, err)
   continue
  }
  for _, x := range ms {
   key := fmt.Sprintf("%d-%d-%d", x.Mid, x.Bid, x.Kid)
   s.merge.SyncAdd(context.Background(), key, x)
  }
  err := msg.Commit()
  log.Info("serviceConsumeproc key:%s partition:%d offset:%d err: %+v, len(%v)", msg.Key, msg.Partition, msg.Offset, err, len(ms))
 }
}

func (s *Service) serviceFlush(merges []*model.Merge) {
 // 相同的mid聚合在一起
 sort.Slice(merges, func(i, j int) bool { return merges[i].Mid < merges[j].Mid })
 var ms []*model.Merge
 for _, m := range merges {
  if (len(ms) < s.c.Job.ServiceBatch) || (ms[len(ms)-1].Mid == m.Mid) {
   ms = append(ms, m)
   continue
  }
  s.FlushCache(context.Background(), ms)
  ms = []*model.Merge{m}
 }
 if len(ms) > 0 {
  s.FlushCache(context.Background(), ms)
 }
}

// FlushCache  数据从缓存写入到DB中
func (s *Service) FlushCache(c context.Context, merges []*model.Merge) (err error) {
 var histories []*model.History
 if histories, err = s.dao.HistoriesCache(c, merges); err != nil {
  log.Error("historyDao.Cache(%+v) error(%v)", merges, err)
  return
 }
 prom.BusinessInfoCount.Add("histories-db", int64(len(histories)))
 if err = s.limit.WaitN(context.Background(), len(histories)); err != nil {
  log.Error("s.limit.WaitN(%v) err: %+v", len(histories), err)
 }
 // 一直尝试到成功
 for {
  if err = s.dao.AddHistories(c, histories); err != nil {
   prom.BusinessInfoCount.Add("retry", int64(len(histories)))
   time.Sleep(time.Duration(s.c.Job.RetryTime))
   continue
  }
  break
 }
 // pipeline的.Do去执行,新建一个chan
 s.cache.Do(c, func(c context.Context) {
  for _, merge := range merges {
   limit := s.c.Job.CacheLen
   s.dao.TrimCache(context.Background(), merge.Business, merge.Mid, limit)
  }
 })
 return
}

func (s *Service) initMerge() {
 s.merge = pipeline.NewPipeline(s.c.Merge)
 s.merge.Split = func(a string) int {
  midStr := strings.Split(a, "-")[0]
  return int(crc32.ChecksumIEEE([]byte(midStr)))
 }
 s.merge.Do = func(c context.Context, ch int, values map[string][]interface{}) {
  var merges []*model.Merge
  for _, vs := range values {
   var t int64
   var m *model.Merge
   for _, v := range vs {
    prom.BusinessInfoCount.Incr("dbus-msg")
    if v.(*model.Merge).Time >= t {
     m = v.(*model.Merge)
    }
   }
   if m.Mid%1000 == 0 {
    log.Info("debug: merge mid:%v, ch:%v, value:%+v", m.Mid, ch, m)
   }
   merges = append(merges, m)
  }
  prom.BusinessInfoCount.Add(fmt.Sprintf("ch-%v", ch), int64(len(merges)))
  s.serviceFlush(merges)
 }
 s.merge.Start()
}
```

# Sales promotion service: Read heavy

Three points: staggering, rate limiting, peak shaving, are all aimed at reducing QPS (Queries Per Second). Optimize the entire system.

## Client-side

Simplify sales promotion content.

## Server-side

Use asynchronous processing with compensation to ensure eventual data consistency. Reduce complex logic and precalculate where possible.

## staggering traffic

## Server-side staggering

If the client's request action is triggered by the server actively, such as server-initiated requests, the server can send batches of requests to the client (assuming that the latency from the server to the client is the same for all).

### Client-side staggering

The commonly used client-side staggering strategy employs a random algorithm. Assuming that all clients receive the time to trigger the request action, the client introduces random delays so that the request action is triggered at time t, where t is randomly distributed within the time interval [0, T]. `t = t0 + random() * T` This ensures that t is uniformly distributed within the interval [0, T]. Theoretically, the new maximum QPS is equal to the old maximum QPS divided by T.

## Rate Limiting Strategy

Rate limiting is typically implemented on the server side, such as using Mtop to return corresponding rate limiting error codes after rate limiting. The following mainly focuses on client-side rate limiting. The goal of client-side rate limiting is to eliminate frequent or unnecessary requests.

- If the user has reached the maximum number of requests, no further requests will be made.
- If the user has been rate limited, it will delay making the second request.
- If the user has obtained the desired return result, no further requests will be made.
- Only allow a certain number of requests after a purchase, rejecting subsequent ones.

## Peak Shaving Strategy

### Set the minimum interval between 2 requests as the minimum effective time interval

### Probabilistic Request

Increase the probability of a user's request being sent by introducing a probability p. `qps = Q / t * p`

### Fairness Strategy

Random algorithm + interpolation algorithm to generate a sequence of valid requests. The probability of a valid request for each user within one activity cycle is P, for example, with a probability of 0.2, meaning 1 out of 5 or 2 out of 10 requests have a chance to be sent.

## Other General Strategies

### Cache

Use BFF (Backend for Frontend) caching middleware

### Multiple Read DBs

Read from multiple read databases using binlog and CDC (Change Data Capture).

# business service layout

## business code +kit code

The basic library kit is an independent project. It is suggested that there should be only one at the company level. Dividing it according to functional directories will bring a lot of management work, so it is recommended to merge and integrate. Specific projects are composed of biz server + kit.

## Scaffolding,New Business Project Code

Use Python Cookiecutter to generate empty project code.  

## biz server module

The app service types in microservices are divided into four categories: interface, service, job, admin.

### api

The protobuf file and api repo are saved consistently, just the directories are different
_API_ is the protocol definition directory, xxxapi.proto protobuf file, and the generated _go_ file. We usually describe the _api_ document directly in the _proto_ file.

### interface

It's an external BFF service that assembles data and accepts requests from users, such as exposing an HTTP interface. The service side is the user.

```go
func (fi *FeedInterface) GetPostById(c *gin.Context) {
 var (
  id  int64
  err error
 )
 strId := c.Param("id")
 if id, err = strconv.ParseInt(strId, 10, 64); err != nil {
  xlog.Errorf("GetPostById ParseInt id", err)
  resp.JSONErr(c, xgin.StatusGinBindError)
  return
 }
 req := &feedv1.GetPostByIdReq{
  Id: id,
 }
 err = req.Validate()
 if err != nil {
  xlog.Errorf("GetPostById validate err", err)
  resp.JSONErr(c, xgin.StatusHttpRequestValidateError)
  return
 }
 m, err := fi.feedbiz.GetPost(c.Request.Context(), req.Id)
 if err != nil {
  xlog.Errorf("FeedInterface GetPost err", err)
  resp.JSONErr(c, xgin.StatusInternalServerError)
  return
 }
 post := m.ToInterfaceProto()
 info, err := fi.userbiz.GetUserInfoByUid(c.Request.Context(), m.Uid)
 if err != nil {
  xlog.Errorf("FeedInterface GetPost GetUserProfileByUid err", err)
  resp.JSONErr(c, xgin.StatusInternalServerError)
  return
 }
 post.AuthorInfo = info.ToProto()
 resPB := &feedv1.GetPostByIdResponse{
  Data: &feedv1.GetPostByIdResponse_Data{
   Post: post,
  },
 }
 resp.JSONSuccess(c, resPB.Data)
}
```

#### Response aligns with protobuf's xxxResponse

swagger+validate+http body+response, it is expected to improve the data as a fixed type later.
That is, we need to manually align gin's response.

```protobuf
response:{
 code
 message
 data:
}

message CreatePostResponse {
 int32 code = 1;
 string message = 2;
 message Data{
  string id = 1 ;
 }
 Data data = 3;
}
```

### Service

Internal microservices, only accept requests from other internal services or gateways, for example, it exposes a gRPC interface only to internal services

### Admin

Operations side and to b side user
Different from _service_, it is more of a service for the operational side, usually with higher data permissions, and isolation brings better code-level security.

### Job

A service for stream task processing, the upstream generally depends on the message broker

### Task

Timing tasks, similar to cronjob, are deployed on the _task_ hosting platform.

## an example

### /cmd

The application directory is responsible for the program: start, shut down, configuration initialization, etc.

### /test

Additional external test application and test data. We can construct the `/test` directory at any time as required.

 For larger projects, having a data subdirectory. For example, we can use `/test/data` or `/test/testdata` (if we need to ignore the content of the directory). Note that Go will also ignore directories or files that start with "." or "_", so there is greater flexibility in how to name the testData directory.

### /model in interface service

Overlap with entity's embed part, use entity's
Provide transformation

#### Convert

Come with conversion
Simply provide model.func for pb<->model<->entity

```go
package model

import (
 mallv1 "redbook/api/domain/mall/v1"
 "redbook/app/domain/mall/model/entity"
 "redbook/common/basemodel"
)

type Sku struct {
 Index     int32    `json:"index,omitempty" bson:"index,omitempty"`
 Name      string   `json:"name,omitempty" bson:"name,omitempty"`
 ProductId int64    `json:"product_id,omitempty" bson:"product_id,omitempty"`
 Product   *Product `json:"product,omitempty" bson:"product,omitempty"`
}
type Skus []*Sku

type SkuOnlyId struct {
 Index     int32  `json:"index,omitempty" bson:"index,omitempty"`
 Name      string `json:"name,omitempty" bson:"name,omitempty"`
 ProductId int64  `json:"product_id,omitempty" bson:"product_id,omitempty"`
}
type SkusOnlyId []*SkuOnlyId

type Collection struct {
 Id_       string           `json:"_id,omitempty" bson:"_id,omitempty"`
 Id        int64            `json:"id,omitempty" bson:"id,omitempty"`
 ShopId    int64            `json:"shop_id,omitempty" bson:"shop_id,omitempty"`
 Name      string           `json:"name,omitempty" bson:"name,omitempty"`
 Title     string           `json:"title,omitempty" bson:"title,omitempty"`
 Cover     *basemodel.Image `json:"cover,omitempty" bson:"cover,omitempty"`
 Video     *basemodel.Video `json:"video,omitempty" bson:"video,omitempty"`
 Tags      *basemodel.Tags  `json:"tags,omitempty" bson:"tags,omitempty"`
 State     int32            `json:"state,omitempty" bson:"state,omitempty"`
 CreatedAt int64            `json:"created_at,omitempty" bson:"created_at,omitempty"`
 UpdatedAt int64            `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
 PublishAt int64            `json:"publish_at,omitempty" bson:"publish_at,omitempty"`
 Skus      *Skus            `json:"skus,omitempty" bson:"skus,omitempty"`
 Version   int32            `json:"version,omitempty" bson:"version,omitempty"`
}
type Collections []*Collection

type CollectionCard struct {
 Id_     string           `json:"_id,omitempty" bson:"_id,omitempty"`
 Id      int64            `json:"id,omitempty"`
 ShopId  int64            `json:"shop_id,omitempty"`
 Name    string           `json:"name,omitempty" bson:"name,omitempty"`
 Title   string           `json:"title,omitempty" bson:"title,omitempty"`
 Cover   *basemodel.Image `json:"cover,omitempty" bson:"cover,omitempty"`
 Video   *basemodel.Video `json:"video,omitempty" bson:"video,omitempty"`
 Tags    *basemodel.Tags  `json:"tags,omitempty"`
 State   int32            `json:"state,omitempty" bson:"state,omitempty"`
 Version int32            `json:"version,omitempty"`
 // Product is Skus[0]
 ProductId   int64        `json:"product_id,omitempty"`
 ProductCard *ProductCard `json:"product_card,omitempty"`
}

type CollectionCards []*CollectionCard

type CreateCollection struct {
 ShopId     int64            `json:"shop_id,omitempty" bson:"shop_id,omitempty"`
 Name       string           `json:"name,omitempty" bson:"name,omitempty"`
 Title      string           `json:"title,omitempty" bson:"title,omitempty"`
 Cover      *basemodel.Image `json:"cover,omitempty" bson:"cover,omitempty"`
 Video      *basemodel.Video `json:"video,omitempty" bson:"video,omitempty"`
 Tags       *basemodel.Tags  `json:"tags,omitempty" bson:"tags,omitempty"`
 PublishAt  int64            `json:"publish_at,omitempty" bson:"publish_at,omitempty"`
 SkusOnlyId *SkusOnlyId      `json:"skus_only_id,omitempty" bson:"skus_only_id,omitempty"`
}

func (m *CollectionCard) ToProto() *mallv1.CollectionCard {
 if m.ProductCard == nil {
  return &mallv1.CollectionCard{
   Id:        m.Id,
   ShopId:    m.ShopId,
   Name:      m.Name,
   Title:     m.Title,
   Cover:     m.Cover.ToProto(),
   Video:     m.Video.ToProto(),
   Tags:      m.Tags.ListToProto(),
   State:     m.State,
   Version:   m.Version,
   ProductId: m.ProductId,
  }
 }
 return &mallv1.CollectionCard{
  Id:          m.Id,
  ShopId:      m.ShopId,
  Name:        m.Name,
  Title:       m.Title,
  Cover:       m.Cover.ToProto(),
  Video:       m.Video.ToProto(),
  Tags:        m.Tags.ListToProto(),
  State:       m.State,
  Version:     m.Version,
  ProductId:   m.ProductId,
  ProductCard: m.ProductCard.ToProto(),
 }
}

func (m *Sku) FromSkuOnlyId(sku *entity.SkuOnlyId) {
 m.Index = sku.Index
 m.Name = sku.Name
 m.ProductId = sku.ProductId
}

func (ms *Skus) ListFromSkuOnlyId(skus *entity.SkusOnlyId) {
 for _, sku := range *skus {
  m := &Sku{}
  m.FromSkuOnlyId(sku)
  *ms = append(*ms, m)
 }
}

func (m *Sku) ToEntitySkuOnlyId() *entity.SkuOnlyId {
 return &entity.SkuOnlyId{
  Index:     m.Index,
  Name:      m.Name,
  ProductId: m.ProductId,
 }
}

func (ms *Skus) ListToEntitySkuOnlyId() *entity.SkusOnlyId {
 skus := make(entity.SkusOnlyId, 0, len(*ms))
 for _, m := range *ms {
  skus = append(skus, m.ToEntitySkuOnlyId())
 }
 return &skus
}

func (m *SkuOnlyId) ToEntity() *entity.SkuOnlyId {
 return &entity.SkuOnlyId{
  Index:     m.Index,
  Name:      m.Name,
  ProductId: m.ProductId,
 }
}
func (m *SkuOnlyId) FromProto(pb *mallv1.SkuOnlyId) {
 m.Index = pb.Index
 m.Name = pb.Name
 m.ProductId = pb.ProductId
}

func (ms *SkusOnlyId) ListToEntity() *entity.SkusOnlyId {
 skus := make(entity.SkusOnlyId, 0, len(*ms))
 for _, m := range *ms {
  m.ToEntity()
  skus = append(skus, m.ToEntity())
 }
 return &skus
}

func (ms *SkusOnlyId) ListFromProto(pbs []*mallv1.SkuOnlyId) {
 for _, pb := range pbs {
  m := &SkuOnlyId{}
  m.FromProto(pb)
  *ms = append(*ms, m)
 }
}

func (m *Sku) FromProto(pb *mallv1.Sku) {
 m.Index = pb.Index
 m.Name = pb.Name
 m.ProductId = pb.ProductId
}

func (ms *Skus) ListFromProto(pbs []*mallv1.Sku) {
 for _, pb := range pbs {
  m := &Sku{}
  m.FromProto(pb)
  *ms = append(*ms, m)
 }
}

func (m *Sku) ToProto() *mallv1.Sku {
 if m.Product == nil {
  return &mallv1.Sku{
   Index:     m.Index,
   Name:      m.Name,
   ProductId: m.ProductId,
  }
 }
 return &mallv1.Sku{
  Index:     m.Index,
  Name:      m.Name,
  ProductId: m.ProductId,
  Product:   m.Product.ToProto(),
 }
}

func (ms *Skus) ListToProto() []*mallv1.Sku {
 pbs := make([]*mallv1.Sku, 0, len(*ms))
 for _, m := range *ms {
  pbs = append(pbs, m.ToProto())
 }
 return pbs
}

// ToCard convert Collection to CollectionCard
// but without ProductCard
func (m *Collection) ToCard() *CollectionCard {
 skus := *m.Skus
 productId := skus[0].ProductId
 return &CollectionCard{
  Id:        m.Id,
  ShopId:    m.ShopId,
  Name:      m.Name,
  Title:     m.Title,
  Cover:     m.Cover,
  Video:     m.Video,
  Tags:      m.Tags,
  State:     m.State,
  Version:   m.Version,
  ProductId: productId,
 }
}

func (m *Collection) ToEntity() (*entity.Collection, error) {
 skus := m.Skus
 return &entity.Collection{
  Id:         m.Id,
  ShopId:     m.ShopId,
  Name:       m.Name,
  Title:      m.Title,
  Cover:      m.Cover,
  Video:      m.Video,
  Tags:       m.Tags,
  State:      m.State,
  CreatedAt:  m.CreatedAt,
  UpdatedAt:  m.UpdatedAt,
  PublishAt:  m.PublishAt,
  Version:    m.Version,
  SkusOnlyId: skus.ListToEntitySkuOnlyId(),
 }, nil
}

func (ms *Collections) ListToEntity() ([]*entity.Collection, error) {
 ets := make([]*entity.Collection, 0, len(*ms))
 for _, m := range *ms {
  et, err := m.ToEntity()
  if err != nil {
   return nil, err
  }
  ets = append(ets, et)
 }
 return ets, nil
}

func (m *Collection) FromEntity(et *entity.Collection) error {
 m.Id_ = et.Id_.Hex()
 m.Id = et.Id
 m.ShopId = et.ShopId
 m.Name = et.Name
 m.Title = et.Title
 m.Cover = et.Cover
 m.Video = et.Video
 m.Tags = et.Tags
 m.State = et.State
 m.CreatedAt = et.CreatedAt
 m.UpdatedAt = et.UpdatedAt
 m.PublishAt = et.PublishAt
 m.Version = et.Version
 m.Skus = &Skus{}
 m.Skus.ListFromSkuOnlyId(et.SkusOnlyId)
 return nil
}
func (ms *Collections) ListFromEntity(ets []*entity.Collection) error {
 for _, et := range ets {
  m := &Collection{}
  if err := m.FromEntity(et); err != nil {
   return err
  }
  *ms = append(*ms, m)
 }
 return nil
}

func (ms *Collections) MapFromEntity(ets map[int64]*entity.Collection) error {
 for _, et := range ets {
  m := &Collection{}
  if err := m.FromEntity(et); err != nil {
   return err
  }
  *ms = append(*ms, m)
 }
 return nil
}

func (m *Collection) ToProto() *mallv1.Collection {
 return &mallv1.Collection{
  Id:        m.Id,
  ShopId:    m.ShopId,
  Name:      m.Name,
  Title:     m.Title,
  Cover:     m.Cover.ToProto(),
  Video:     m.Video.ToProto(),
  Tags:      m.Tags.ListToProto(),
  State:     m.State,
  CreatedAt: m.CreatedAt,
  UpdatedAt: m.UpdatedAt,
  PublishAt: m.PublishAt,
  Version:   m.Version,
  Skus:      m.Skus.ListToProto(),
 }
}

func (ms *Collections) ListToProto() []*mallv1.Collection {
 pbs := make([]*mallv1.Collection, 0, len(*ms))
 for _, m := range *ms {
  pbs = append(pbs, m.ToProto())
 }
 return pbs
}

func (m *Collection) FromProto(pb *mallv1.Collection) {
 m.Id = pb.Id
 m.ShopId = pb.ShopId
 m.Name = pb.Name
 m.Title = pb.Title
 m.Cover = &basemodel.Image{}
 m.Cover.FromProto(pb.Cover)
 m.Video = &basemodel.Video{}
 m.Video.FromProto(pb.Video)
 m.Tags = &basemodel.Tags{}
 m.Tags.ListFromProto(pb.Tags)
 m.Skus = &Skus{}
 m.Skus.ListFromProto(pb.Skus)
}

func (m *CreateCollection) FromProto(pb *mallv1.CreateCollection) {
 m.ShopId = pb.ShopId
 m.Name = pb.Name
 m.Title = pb.Title
 m.Cover = &basemodel.Image{}
 m.Cover.FromProto(pb.Cover)
 m.Video = &basemodel.Video{}
 m.Video.FromProto(pb.Video)
 m.Tags = &basemodel.Tags{}
 m.Tags.ListFromProto(pb.Tags)
 m.PublishAt = pb.PublishAt
 m.SkusOnlyId = &SkusOnlyId{}
 m.SkusOnlyId.ListFromProto(pb.SkusOnlyId)
}

```

### /internal/data

### /internal/data/entity

- mongo,struct
- Kafka protobuf
- grpcclient protobuf
- redis json or protobuf
- ...

### /model in domain service

for business layer use
also Provide const

```go
const (
 ChanStateOffline = int32(0) //频道下线
 ChanStateStop = int32(1) //频道停用
 ChanStateCommon = int32(2) //频道普通
 ChanStateRecomend = int32(3) //频道推荐
)
```
