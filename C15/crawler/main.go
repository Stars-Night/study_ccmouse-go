package main

import (
	"ccmouse-go/C15/crawler/engine"
	"ccmouse-go/C15/crawler/scheduler"
	"ccmouse-go/C15/crawler/zhenai/parser"
)

//多协程调度器（scheduler）
//多协程Worker
func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}
