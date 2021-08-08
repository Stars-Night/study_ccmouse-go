package main

import (
	"ccmouse-go/C16/crawler4/engine"
	"ccmouse-go/C16/crawler4/persist"
	"ccmouse-go/C16/crawler4/scheduler"
	"ccmouse-go/C16/crawler4/zhenai/parser"
)

// 前端模板演示
func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}
