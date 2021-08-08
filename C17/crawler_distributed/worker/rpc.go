package worker

import (
	"ccmouse-go/C17/crawler/engine"
)

// 实现RPC
type CrawlService struct{}

// 把客户端传过来的参数反序列化，执行worker函数，返回序列化的结果
func (CrawlService) Process(req Request, result *ParserResult) error {
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}

	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}

	*result = SerializeResult(engineResult)
	return nil
}
