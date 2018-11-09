package scheduler

import "com.yuer.gio/lgotest/simple/crawler/concurrent/engine"

type SimpleScheduler struct {
	workChan chan  engine.Request
}


func (s *SimpleScheduler) ConfigMasterWorkerChan(r  chan engine.Request) {
	 	s.workChan = r
}



func (s *SimpleScheduler) Commit(r engine.Request){
	go func() {
		s.workChan <- r
	}()

}