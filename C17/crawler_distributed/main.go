package main

import (
	"ccmouse-go/C17/crawler/engine"
	"ccmouse-go/C17/crawler/scheduler"
	"ccmouse-go/C17/crawler/zhenai/parser"
	"ccmouse-go/C17/crawler_distributed/config"
	itemsaver "ccmouse-go/C17/crawler_distributed/persist/client"
	"ccmouse-go/C17/crawler_distributed/rpcsupport"
	worker "ccmouse-go/C17/crawler_distributed/worker/client"
	"flag"
	"log"
	"net/rpc"
	"strings"
)

// 用命令行执行该文件的时候可以给itemsaver_host、worker_hosts传参
// 字符串类型，多个参数用逗号分隔
// go run main.go --itemsaver_host=":1234" --worker_hosts=":9000,:9001"
var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts   = flag.String("worker_hosts", "", "worker hosts (comma separated)")
)

// Todo 使用脚本进行部署，把参数定义一下，批处理
func main() {
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))
	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParserCityList, config.ParserCityList),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("Error connecting to %s: %v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		// 死循环发送clients slice里面的rpc客户端给CreateProcessor函数
		// CreateProcessor函数依次接收，这样就可以轮询调用不同的worker服务
		for {
			for _, c := range clients {
				out <- c
			}
		}
	}()
	return out
}
