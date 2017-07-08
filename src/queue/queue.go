package queue

import (
	"net"
	"sync"
)

type Queue struct {
	sync.Mutex
	queue []net.Conn
}

func NewQueue() *Queue {
	return &Queue{
		queue: make([]net.Conn, 0),
	}
}

func (q *Queue) Push(c net.Conn) {
	q.Lock()
	defer q.Unlock()
	q.queue = append(q.queue, c)
	return
}

func (q *Queue) Pop() (c net.Conn) {
	defer q.Unlock()
	q.Lock()
	c = (q.queue)[0]
	q.queue = (q.queue)[1:]
	return
}

func (q *Queue) Len() int {
	defer q.Unlock()
	q.Lock()
	len := len(q.queue)
	return len
}
