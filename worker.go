package worker

import (
	logging "github.com/ipfs/go-log"
)

var log = logging.Logger("go-worker")

type Work interface {
	Run()
}

type Worker struct {
	ID          int
	Work        chan Work
	WorkerQueue chan chan Work
	QuitChan    chan bool
}

func NewWorker(id int, workerQueue chan chan Work) Worker {
	// Create, and return the worker.
	worker := Worker{
		ID:          id,
		Work:        make(chan Work),
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool)}

	return worker
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerQueue <- w.Work

			select {
			case work := <-w.Work:
				log.Infof("worker %d: Received work request", w.ID)
				work.Run()
			case <-w.QuitChan:
				log.Infof("worker %d : stopping", w.ID)
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	w.QuitChan <- true
}
