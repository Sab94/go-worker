package main

import (
	"time"

	gw "github.com/Sab94/go-worker"
	logging "github.com/ipfs/go-log"
)

var log = logging.Logger("go-worker-example")

type Work struct {
	Name string
}

func (w Work) Run() {
	log.Debugf("Executing Task with name %s", w.Name)
}

func main() {
	logging.SetLogLevel("*", "Debug")
	manager := gw.NewManager(2)

	work := Work{
		Name: "Example Work",
	}

	workChannle := manager.NewBufferedManager(5)
	workChannle <- work
	time.Sleep(time.Second * 2)

	manager.GoWork(work)
	time.Sleep(time.Second * 2)
	return
}
