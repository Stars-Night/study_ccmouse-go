package main

import (
	"ccmouse-go/C16/crawler6/engine"
	"ccmouse-go/C16/crawler6/persist"
	"ccmouse-go/C16/crawler6/scheduler"
	"ccmouse-go/C16/crawler6/zhenai/parser"
)

// 完善前端展示
// fill in query string
// support search button
// support paging
// add start page
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
