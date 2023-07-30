package entity

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HistoryItem struct {
	Id_       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Id        int64
	PostID    int64
	Uid       int64
	CreatedAt int64
}

type HistoryItems []*HistoryItem

func (et *HistoryItem) FromCache(v *redis.Z) error {
	score, err := strconv.ParseInt(fmt.Sprintf("%1.0f", v.Score), 10, 64)
	if err != nil {
		return err
	}
	postIdStr := v.Member.(string)

	postId, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {
		return err
	}
	et.CreatedAt = score
	et.PostID = postId
	return nil
}
