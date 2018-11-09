package engine

import "fmt"

//并发爬虫项目

type  ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Commit(Request)
	ConfigMasterWorkerChan(chan Request)
}

var index = 0
func (e *ConcurrentEngine) Run(seeds ... Request){


	in := make(chan Request)
	out := make(chan  ParseResult)
	e.Scheduler.ConfigMasterWorkerChan(in)

	//创建work ，所有的work公用一个输入
	for i:=0 ;i <e.WorkerCount;i++{
		createWorkrer(in,out)
	}

	//最好是等他们做完以后，在进行submit
	for _, r := range  seeds{
		e.Scheduler.Commit(r)
	}

	for{
		result := <- out
		for _,item := range result.Items{
			index ++
			fmt.Printf("Got %d item %v \n",index ,item)
		}
		for _,request :=range result.Requests{
			e.Scheduler.Commit(request)
		}
	}
}

func createWorkrer(in chan Request, out chan ParseResult) {
	go func() {
		for  {
			//request := <-  in

			var request Request
			request = <-  in
			result, err := work(request)
			if err != nil{
				continue
			}
			out  <- result
		}
	}()
}