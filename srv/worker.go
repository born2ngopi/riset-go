package srv

import (
	"log"
	"net"
)

const DEFAULT_WORKER int = 5

type Worker struct {
	WorkerFunc func(conn net.Conn)
	MaxWorker int
	Conn chan net.Conn
}

func (w *Worker) Start() {
	for i := 0; i < w.MaxWorker; i++ {

		go func() {

			for c := range w.Conn {
				log.Println("new connection processed")
				w.WorkerFunc(c)
			}

		}()

	}
}

func (w *Worker) Add(conn net.Conn) {
	w.Conn <- conn
}
