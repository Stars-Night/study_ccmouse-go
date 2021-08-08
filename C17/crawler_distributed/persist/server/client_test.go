package main

import (
	"ccmouse-go/C17/crawler/engine"
	"ccmouse-go/C17/crawler/model"
	"ccmouse-go/C17/crawler_distributed/config"
	"ccmouse-go/C17/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"

	// start ItemSaverServer
	go serverRpc(host, "test1")
	time.Sleep(time.Second)

	// start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// call save
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1106076361",
		Type: "zhenai",
		Id:   "1106076361",
		Payload: model.Profile{
			Age:       23,
			Height:    155,
			Weight:    45,
			Gender:    "女士",
			Income:    "5-8千",
			Marriage:  "未婚",
			Education: "高中及以下",
			Xinzuo:    "魔羯座",
			Car:       "未买车",
			Name:      "萍儿",
			House:     "租房",
			Hokou:     "湖南娄底",
		},
	}

	var result string
	err = client.Call(config.ItemSaverRpc, item, &result)

	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
