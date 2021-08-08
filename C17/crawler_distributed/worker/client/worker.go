package client

import (
	"ccmouse-go/C17/crawler/engine"
	"ccmouse-go/C17/crawler_distributed/config"
	"ccmouse-go/C17/crawler_distributed/worker"
	"net/rpc"
)

/**
此函数创建一个分布式版的worker方法
内部匿名函数：使用RPC调用服务端的worker方法，把序列化后的Request发送过去，把返回的结果反序列化
*/
func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	//client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort0))
	//if err != nil {
	//	return nil, err
	//}

	return func(req engine.Request) (engine.ParserResult, error) {
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParserResult
		c := <-clientChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParserResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}
}
