package main

import (
	"ccmouse-go/C16/crawler3/engine"
	"ccmouse-go/C16/crawler3/persist"
	"ccmouse-go/C16/crawler3/scheduler"
	"ccmouse-go/C16/crawler3/zhenai/parser"
)

// 把每个item的client提到ItemSaver共用一个
// 把city.go中ParserFunc传递的匿名函数提取出来；ParserFunc新增了url参数，由engine.worker传参
// 把worker从simple.go提出到新建的worker.go
// elastic search的index采用参数传递
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
