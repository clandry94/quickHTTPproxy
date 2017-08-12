package worker

import (
	"bufio"
	"crypto/sha1"
	"fmt"
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

func (sw SortingWorker) Run() {
	func() {
		for {

			select {
			case work := <-sw.newConnections:
				logger.Info("New work for worker")
				buffReader := bufio.NewReader(work)
				bytes, err := buffReader.ReadBytes('\n')
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("%s", bytes)
				}
				work.Close()
				logger.Info("Connection closed")
			case <-sw.quit:
				logger.Info("Closing worker", sw.id)
			}

		}
	}()
}
