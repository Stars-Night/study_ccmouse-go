package main

import (
	"ccmouse-go/C17/crawler/engine"
	"ccmouse-go/C17/crawler/persist"
	"ccmouse-go/C17/crawler/scheduler"
	"ccmouse-go/C17/crawler/zhenai/parser"
	"ccmouse-go/C17/crawler_distributed/config"
)

func main() {
	itemChan, err := persist.ItemSaver(config.ElasticIndex)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParserCityList, config.ParserCityList),
	})
}
