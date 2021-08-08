package main

import (
	"ccmouse-go/C16/crawler2/engine"
	"ccmouse-go/C16/crawler2/persist"
	"ccmouse-go/C16/crawler2/scheduler"
	"ccmouse-go/C16/crawler2/zhenai/parser"
)

//把数据存elasticsearch
//添加URL与ID：定义Item struct，属性有Url, Id, Type、model.Profile
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
