all-demo :
# bff/demo
	protoc -I . \
	-I ../ \
	-I ../third_party/proto/ \
	--go_out=../../ \
	--validate_out=lang=go:../../ \
    --openapiv2_out . \
	--openapiv2_opt logtostderr=true \
	./interface/demo/v1/demo.proto
# domain/demo
	protoc -I .  -I ../ \
	-I ../third_party/proto/ \
	--go_out=../../ ./domain/demo/v1/demo.proto \
	--go-grpc_out=../../ ./domain/demo/v1/demo.proto
inject-tag:
	protoc-go-inject-tag -input="./*/*/*.pb.go"
redbookpb-v1:
	protoc -I .  -I ../ \
	-I ../third_party/proto/ \
	--go_out=../../ ./redbookpb/v1/base.proto
	protoc-go-inject-tag -input=./redbookpb/v1/base.pb.go
interface-demo:
	protoc -I . \
	-I ../ \
	-I ../third_party/proto/ \
	--go_out=../../ \
	--validate_out=lang=go:../../ \
    --openapiv2_out . \
	--openapiv2_opt logtostderr=true \
	./interface/demo/v1/demo.proto

domain-user:
# domain/user
	protoc  -I . \
	-I ../ \
	-I ../third_party/proto/ \
	--go_out=../../ ./domain/user/v1/user.proto \
	--go-grpc_out=../../ ./domain/user/v1/user.proto \
	--validate_out=lang=go:../../ ./domain/user/v1/user.proto
interface-comment:
# interface/comment
	protoc -I .  -I ../ \
	-I ../third_party/proto/ \
    --go_out=../../ \
	--validate_out=lang=go:../../ \
    --openapiv2_out . \
	--openapiv2_opt logtostderr=true \
	./interface/comment/v1/comment.proto
	protoc-go-inject-tag -input="./interface/comment/v1/*.pb.go"
domain-comment:
	protoc  -I . \
	-I ../ \
	-I ../third_party/proto/ \
	--go_out=../../ ./domain/comment/v1/comment.proto \
	--go-grpc_out=../../ ./domain/comment/v1/comment.proto \
	--validate_out=lang=go:../../ \
	./domain/comment/v1/comment.proto
job-comment:
	protoc -I .  -I ../ \
	-I ../third_party/proto/ \
    --go_out=../../ \
	--validate_out=lang=go:../../ \
	./job/comment/v1/comment.proto

domain-like:
	protoc -I .  -I ../ \
	-I ../third_party/proto/ \
    --go_out=../../ \
	--go-grpc_out=../../ \
	--validate_out=lang=go:../../ \
	./domain/like/v1/like.proto

interface-feed:
	protoc -I . \
	-I ../ \
	-I ../third_party/proto/ \
	--go_out=../../ \
	--validate_out=lang=go:../../ \
    --openapiv2_out . \
	--openapiv2_opt logtostderr=true \
	./interface/feed/v1/feed.proto
	protoc-go-inject-tag -input="./interface/feed/v1/*.pb.go"

domain-feed:
	protoc -I .  -I ../ \
	-I ../third_party/proto/ \
    --go_out=../../ \
	--go-grpc_out=../../ \
	--validate_out=lang=go:../../ \
	./domain/feed/v1/feed.proto


interface-mall:
	protoc -I . \
	-I ../ \
	-I ../third_party/proto/ \
	--go_out=../../ \
	--validate_out=lang=go:../../ \
    --openapiv2_out . \
	--openapiv2_opt logtostderr=true \
	./interface/mall/v1/mall.proto
	protoc-go-inject-tag -input="./interface/mall/v1/*.pb.go"

domain-mall:
	protoc -I .  -I ../ \
	-I ../third_party/proto/ \
    --go_out=../../ \
	--go-grpc_out=../../ \
	--validate_out=lang=go:../../ \
	./domain/mall/v1/mall.proto
