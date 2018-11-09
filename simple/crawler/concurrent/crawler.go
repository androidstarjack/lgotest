package main

import (
	"com.yuer.gio/lgotest/simple/crawler/concurrent/engine"
	"com.yuer.gio/lgotest/simple/crawler/concurrent/scheduler"
	"com.yuer.gio/lgotest/simple/crawler/concurrent/zhenai/parse"
)

const URL  =  "http://www.zhenai.com/zhenghun"

//并发任务方式进行爬虫
///**
// *
// */
func main() {
	e := engine.ConcurrentEngine{
		Scheduler:  &scheduler.SimpleScheduler{ },
		WorkerCount:10,
	}
	e.Run(engine.Request{
		Url:URL,
		ParseFunc:parse.ParseCityListdata,
	})
}
