package main

import (
	"ccmouse-go/C15/crawler3/engine"
	"ccmouse-go/C15/crawler3/scheduler"
	"ccmouse-go/C15/crawler3/zhenai/parser"
)

//重构Scheduler使得Engine可以共用queued和simple的代码
//把Scheduler interface中的WorkerReady方法拆出来，变成组合接口，使得作为参数传递时更轻
func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		//Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}
