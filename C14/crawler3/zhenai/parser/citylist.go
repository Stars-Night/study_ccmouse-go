package parser // Package parser 爬虫解析包

import (
	"ccmouse-go/C14/crawler3/engine"
	"regexp"
)

// 当前页面要匹配的网址的正则表达式，这些网址是下一步要继续爬取内容的
const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// ParserCityList 城市列表页面的解析器
func ParserCityList(content []byte) engine.ParserResult {
	// 获取正则匹配到的结果
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(content, -1)

	result := engine.ParserResult{} // 定义一个解析结果的结构体，以便保存需要的解析结果
	// 遍历处理匹配结果
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2])) // 保存要获取的内容
		// 下一步要请求的内容
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),     // 下一步要请求的地址
			ParserFunc: engine.NilParser, // 下一步请求内容的解析器，这里是空解析器，下一步不用解析了
		})
	}

	return result
}
