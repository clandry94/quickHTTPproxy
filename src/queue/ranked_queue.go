package queue

import (
	"net"
)

type Tag string
type Priority int

type RankedQueue struct {
	queue    *Queue
	Tag      Tag
	Priority Priority
}

func NewRankedQueue(tag string, priority int) RankedQueue {
	rankedQueue := RankedQueue{
		queue:    NewQueue(),
		Tag:      Tag(tag),
		Priority: Priority(priority),
	}
	return rankedQueue
}

func (rq *RankedQueue) Push(c net.Conn) {
	rq.queue.Push(c)
}

func (rq *RankedQueue) Pop() net.Conn {
	return rq.queue.Pop()
}

func (rq *RankedQueue) Len() int {
	return rq.queue.Len()
}
