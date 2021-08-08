package main

import (
	"ccmouse-go/C16/crawler/engine"
	"ccmouse-go/C16/crawler/persist"
	"ccmouse-go/C16/crawler/scheduler"
	"ccmouse-go/C16/crawler/zhenai/parser"
)

//只生成profile的Items
//在engine里面为每个item开goroutine，把item传给ItemSaver
func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}
