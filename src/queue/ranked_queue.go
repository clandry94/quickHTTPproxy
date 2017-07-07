package queue

type Tag string
type Priority int

type RankedQueue struct {
	queue    *Queue
	Tag      Tag
	Priority Priority
}

func NewRankedQueue(tag string, priority int) RankedQueue {
	glog.Info("Creating ranked queue")
	rankedQueue := RankedQueue{
		queue:    NewQueue(),
		Tag:      Tag(tag),
		Priority: Priority(priority),
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
