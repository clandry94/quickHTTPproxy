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
	glog.Info("Creating new ranked queue map")
	return &RankedQueueMap{
		Id:             sha1.New(),
		RankedQueueMap: make(map[Tag]*RankedQueue),
	}
}

func (rqm *RankedQueueMap) Insert(q *RankedQueue) {
	rqm.RankedQueueMap[q.Tag] = q
}
