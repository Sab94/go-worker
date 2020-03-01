package worker

type Manager struct {
	workerQueue chan chan Work
	workerMap   map[int]*Worker
	workChannel chan Work
	stopSignal  chan bool
	stopped     bool
}

func NewManager(workerCount int) *Manager {
	manager := Manager{
		workerQueue: make(chan chan Work, workerCount),
		workerMap:   make(map[int]*Worker, workerCount),
		stopped:     false,
		stopSignal:  make(chan bool),
	}

	for i := 0; i < workerCount; i++ {
		log.Infof("Starting worker : %d", i+1)
		worker := NewWorker(i+1, manager.workerQueue)
		worker.Start()
		manager.workerMap[i] = &worker
	}

	return &manager
}

func (m *Manager) GoWork(newWork Work) {
	go func() {
		worker := <-m.workerQueue
		log.Info("Dispatching work request")
		worker <- newWork
	}()
}

func (m *Manager) StopWork() {
	log.Info("Got stop request")
	go func() {
		log.Info("Stopping workers")
		for i := 1; i <= len(m.workerMap); i++ {
			m.workerMap[i].Stop()
		}
		m.stopped = true
		m.stopSignal <- true
	}()
}

func (m *Manager) NewBufferedManager(size int) chan Work {
	if size > 1000 {
		log.Error("Queue size too big")
		return nil
	}
	m.workChannel = make(chan Work, size)
	m.startBufferedManager()
	return m.workChannel
}

func (m *Manager) startBufferedManager() {
	go func() {
		log.Info("Starting listener for Tasks")
		for {
			select {
			case work := <-m.workChannel:
				log.Info("Got new task on channel")
				go func() {
					worker := <-m.workerQueue
					log.Info("Assigning worker")
					worker <- work
				}()
			case <-m.stopSignal:
				return
			}
		}
	}()
}
