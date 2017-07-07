package queue

import (
	"sync"
)

type Queue struct {
	sync.Mutex
	queue []string
}

func NewQueue() *Queue {
	return &Queue{
		queue: make([]string, 0),
	}
}

func (q *Queue) Push(s string) {
	q.Lock()
	defer q.Unlock()
	q.queue = append(q.queue, s)
	return
}

func (q *Queue) Pop() (s string) {
	defer q.Unlock()
	q.Lock()
	s = (q.queue)[0]
	q.queue = (q.queue)[1:]
	return
}

func (q *Queue) Len() int {
	defer q.Unlock()
	q.Lock()
	len := len(q.queue)
	return len
}
