package engine

import (
	"ccmouse-go/C16/crawler4/fetcher"
	"log"
)

func worker(request Request) (ParserResult, error) {
	body, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("Fetch error: %v, url: %s:", request.Url, err)
		return ParserResult{}, err
	}

	return request.ParserFunc(body, request.Url), nil
}
