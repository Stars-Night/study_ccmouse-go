package parser

import (
	"ccmouse-go/C14/crawler4/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParserCityList(content []byte) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(content, -1)

	result := engine.ParserResult{}
	limit := 5

	for _, m := range matches {
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParserCity,
		})
		limit--
		if limit == 0 {
			break
		}
	}

	return result
}