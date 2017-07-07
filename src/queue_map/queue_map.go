package queue_map

import (
	"crypto/sha1"
	"github.com/clandry94/quickHTTPproxy/src/queue"
	"hash"
)

type RankedQueueMap struct {
	Id             hash.Hash
	RankedQueueMap map[queue.Tag]*queue.RankedQueue
}

func NewRankedQueueMap() *RankedQueueMap {
	return &RankedQueueMap{
		Id:             sha1.New(),
		RankedQueueMap: make(map[queue.Tag]*queue.RankedQueue),
	}
}
