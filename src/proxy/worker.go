package proxy

import (
	"crypto/sha1"
	"fmt"
	"github.com/clandry94/quickHTTPproxy/src/queue"
	"hash"
)

func NewConnWorkerPool(size int, rankedQueueMap *queue.RankedQueueMap) {

	for i := 0; i < size; i++ {
		worker := connWorker{
			ID:             createWorkerID(),
			RankedQueueMap: rankedQueueMap,
			Status:         "stopped",
			stop:           make(chan bool, 1),
		}
		go worker.Run()
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
					fmt.Println(conn)
					//conn.Close()
				}
			}
		}
	}()
}
