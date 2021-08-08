package main

import (
	"ccmouse-go/C16/crawler5/engine"
	"ccmouse-go/C16/crawler5/persist"
	"ccmouse-go/C16/crawler5/scheduler"
	"ccmouse-go/C16/crawler5/zhenai/parser"
)

// MVC结构展示数据
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
