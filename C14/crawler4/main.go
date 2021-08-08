package main

import (
	"ccmouse-go/C14/crawler4/engine"
	"ccmouse-go/C14/crawler4/zhenai/parser"
)

//对应14-7、14-8、14-9课程
func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}
