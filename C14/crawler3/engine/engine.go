package engine // Package engine 爬虫引擎包

import (
	"ccmouse-go/C14/crawler3/fetcher"
	"log"
)

func Run(seeds ...Request) {
	// 获取当前方法的参数
	var requests []Request
	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	// 遍历处理参数，即循环执行要抓取的网页
	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		log.Printf("Fetching Url: %s", request.Url)
		body, err := fetcher.Fetch(request.Url) // 抓取url的内容
		if err != nil {
			log.Printf("Fetch error: %v, url: %s:", request.Url, err)
			continue
		}

		parserResult := request.ParserFunc(body)              // 解析被抓取的内容
		requests = append(requests, parserResult.Requests...) // 把下一步要执行的请求结构体添加到当前请求结构体切片
		// 打印解析的数据（这个简易版爬虫并未把数据保存到DB）
		for _, item := range parserResult.Items {
			log.Printf("got item: %v", item)
		}
	}
}
