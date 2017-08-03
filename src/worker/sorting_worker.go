package worker

import (
	"crypto/sha1"
	"github.com/ivahaev/go-logger"
	"hash"
	"net"
)

type SortingWorker struct {
	id             hash.Hash
	newConnections chan net.Conn
	quit           chan bool
}

func NewSortingWorker(buf chan net.Conn) SortingWorker {
	return SortingWorker{
		id:             sha1.New(),
		newConnections: buf,
		quit:           make(chan bool),
	}
}

func (sw SortingWorker) stop() {
	go func() {
		sw.quit <- true
	}()
}

func (sw SortingWorker) run() {
	func() {
		for {

			select {
			case work := <-sw.newConnections:
				logger.Info("New work for worker %v", sw.id)
				work.Close()
			case <-sw.quit:
				logger.Info("Closing worker %v", sw.id)
			}

		}
	}()
}
