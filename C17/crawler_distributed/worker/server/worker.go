package main

import (
	"ccmouse-go/C17/crawler_distributed/rpcsupport"
	"ccmouse-go/C17/crawler_distributed/worker"
	"flag"
	"fmt"
	"log"
)

// 用命令行执行该文件的时候可以给port传参
// go run worker.go --port=9000
var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServerRpc(fmt.Sprintf(":%d", *port), worker.CrawlService{}))
}
