package main

import (
	"ccmouse-go/C15/crawler4/engine"
	"ccmouse-go/C15/crawler4/scheduler"
	"ccmouse-go/C15/crawler4/zhenai/parser"
)

//爬取城市的所有页面
//Url去重
//用户页面的猜你喜欢做了防爬机制，不爬
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
