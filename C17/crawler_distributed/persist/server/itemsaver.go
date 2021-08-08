package main

import (
	"ccmouse-go/C17/crawler_distributed/config"
	"ccmouse-go/C17/crawler_distributed/persist"
	"ccmouse-go/C17/crawler_distributed/rpcsupport"
	"flag"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

// 用命令行执行该文件的时候可以给port传参
// go run itemsaver.go --port=1234
var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(serverRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func serverRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.99.100:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		return err
	}

	err = rpcsupport.ServerRpc(host, &persist.ItemSaverService{client, index})
	return err
}
