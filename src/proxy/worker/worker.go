package worker

import (
	"crypto/sha1"
	"github.com/clandry94/quickHTTPproxy/src/queue"
	"hash"
)

type ConnWorkerPool chan *connWorker

func NewConnWorkerPool(size int, rankedQueueMap *queue.RankedQueueMap) *ConnWorkerPool {
	connWorkerPool := make(ConnWorkerPool, size)

	for i := 0; i < size; i++ {
		connWorkerPool <- &connWorker{
			ID:             createWorkerID(),
			RankedQueueMap: rankedQueueMap,
			Status:         "stopped",
			stop:           make(chan bool, 1),
		}
	}

	return &connWorkerPool
}

func (cwp *ConnWorkerPool) Start() {
	for {
		worker := <-cwp
		worker.Run()
		cwp <- worker
	}
}

func createWorkerID() hash.Hash {
	return sha1.New()
}

type connWorker struct {
	ID             hash.Hash
	RankedQueueMap *queue.RankedQueueMap
	Status         string
	stop           chan bool
}

func (cw *connWorker) Run() {
	go func() {
		cw.Status = "running"
		for {
			select {
			case <-cw.stop:
				return
			default:
				if cw.RankedQueueMap.RankedQueueMap["conor"].Len() > 0 {
					conn := cw.RankedQueueMap.RankedQueueMap["conor"].Pop()
					conn.Close()
				}
			}
		}
	}()
}
