package queue

type Tag string
type Priority int

type RankedQueue struct {
	queue    Queue
	Tag      Tag
	Priority Priority
}

func NewRankedQueue(q Queue, t string, p int) RankedQueue {
	rankedQueue := RankedQueue{
		queue:    q,
		Tag:      Tag(t),
		Priority: Priority(p),
	}
	return rankedQueue
}

func (rq *RankedQueue) Push(s string) {
	rq.queue.Push(s)
}

func (rq *RankedQueue) Pop() string {
	return rq.queue.Pop()
}

func (rq *RankedQueue) Len() int {
	return rq.queue.Len()
}
