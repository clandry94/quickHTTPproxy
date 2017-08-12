package worker

import (
	"bufio"
	"crypto/sha1"
	"github.com/ivahaev/go-logger"
	"hash"
	"net"
	"net/http"
)

type SortingWorker struct {
	id             hash.Hash
	newConnections chan net.Conn
	Requests       chan *http.Request
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

func (sw SortingWorker) Run() {
	func() {
		for {

			select {
			case work := <-sw.newConnections:
				logger.Info("New work for worker")
				buffReader := bufio.NewReader(work)
				err := sw.pushHttpRequest(buffReader)
				if err != nil {
					logger.Error(err)
				}
				//logger.Info(req)
				//work.Close()
				//logger.Info("Connection closed")
			case <-sw.quit:
				logger.Info("Closing worker", sw.id)
			}

		}
	}()
}

func (sw SortingWorker) pushHttpRequest(b *bufio.Reader) error {
	req, err := http.ReadRequest(b)
	if err != nil {
		return err
	}
	sw.Requests <- req

	// add some debug logging
	return err
}
