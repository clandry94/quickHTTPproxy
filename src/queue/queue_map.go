package queue

type QueueMap map[Tag]RankedQueue

func NewQueueMap() QueueMap {
	qm := make(QueueMap)
	return qm
}
