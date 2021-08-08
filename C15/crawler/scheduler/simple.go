package scheduler

import "ccmouse-go/C15/crawler/engine"

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.WorkerChan <- r }()
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(in chan engine.Request) {
	s.WorkerChan = in
}
