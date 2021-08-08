package engine

import (
	"ccmouse-go/C17/crawler/fetcher"
	"log"
)

func Worker(request Request) (ParserResult, error) {
	body, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("Fetch error: %v, url: %s:", request.Url, err)
		return ParserResult{}, err
	}

	return request.Parser.Parser(body, request.Url), nil
}
