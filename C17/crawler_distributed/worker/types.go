package worker

import (
	"ccmouse-go/C17/crawler/engine"
	"ccmouse-go/C17/crawler/zhenai/parser"
	"ccmouse-go/C17/crawler_distributed/config"
	"fmt"
	"github.com/pkg/errors"
	"log"
)

// 函数序列化的容器， Name，Args都能在网上传递
type SerializedParser struct {
	Name string
	Args interface{}
}

// 原来的Request有接口不能在网络上传输，重新定义
type Request struct {
	Url    string
	Parser SerializedParser
}

type ParserResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParserResult) ParserResult {
	result := ParserResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, nil
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

func DeserializeResult(r ParserResult) engine.ParserResult {
	result := engine.ParserResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}

// 根据Parser中的函数名和参数，返回对应的函数
func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParserCityList:
		return engine.NewFuncParser(parser.ParserCityList, config.ParserCityList), nil
	case config.ParserCity:
		return engine.NewFuncParser(parser.ParserCity, config.ParserCity), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParserProfile:
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParser(userName), nil
		} else {
			return nil, fmt.Errorf("invalid args: %v", p.Args)
		}
	default:
		return nil, errors.New("unknown parser name")
	}
}
