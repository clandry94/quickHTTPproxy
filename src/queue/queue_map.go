package queue

import (
	"crypto/sha1"
	"hash"
)

type RankedQueueMap struct {
	Id             hash.Hash
	RankedQueueMap map[Tag]*RankedQueue
}

func NewRankedQueueMap() *RankedQueueMap {
	return &RankedQueueMap{
		Id:             sha1.New(),
		RankedQueueMap: make(map[Tag]*RankedQueue),
	}
}

func (rqm *RankedQueueMap) Insert(q *RankedQueue) {
	rqm.RankedQueueMap[q.Tag] = q
}
