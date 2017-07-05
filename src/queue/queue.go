package queue

import (
	"sync"
)

type Queue []string

var mutex = &sync.Mutex{}

func (q *Queue) Push(s string) {
	mutex.Lock()
	defer mutex.Unlock()
	*q = append(*q, s)
	return
}

func (q *Queue) Pop() (s string) {
	defer mutex.Unlock()
	mutex.Lock()
	s = (*q)[0]
	*q = (*q)[1:]
	return
}

func (q *Queue) Len() int {
	defer mutex.Unlock()
	mutex.Lock()
	len := len(*q)
	return len
}
