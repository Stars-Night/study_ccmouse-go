package main

import (
	"ccmouse-go/C17/crawler_distributed/config"
	"ccmouse-go/C17/crawler_distributed/rpcsupport"
	"ccmouse-go/C17/crawler_distributed/worker"
	"fmt"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServerRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://album.zhenai.com/u/1106076361",
		Parser: worker.SerializedParser{
			Name: config.ParserProfile,
			Args: "萍儿",
		},
	}

	var result worker.ParserResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
