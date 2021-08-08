package main

import (
	"ccmouse-go/C14/crawler3/engine"
	"ccmouse-go/C14/crawler3/zhenai/parser"
)

//对应14-5课程
func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}
