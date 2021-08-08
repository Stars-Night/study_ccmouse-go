package main

import (
	"ccmouse-go/C15/crawler2/engine"
	"ccmouse-go/C15/crawler2/scheduler"
	"ccmouse-go/C15/crawler2/zhenai/parser"
)

//调度器（scheduler）: Request队列、Worker队列
//多协程Worker
func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}
