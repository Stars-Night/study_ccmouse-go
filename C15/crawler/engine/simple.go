package engine

import (
	"ccmouse-go/C15/crawler/fetcher"
	"log"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		parserResult, err := worker(request)
		if err != nil {
			continue
		}

		requests = append(requests, parserResult.Requests...)
		for _, item := range parserResult.Items {
			log.Printf("got item: %v", item)
		}
	}
}

func worker(request Request) (ParserResult, error) {
	log.Printf("Fetching Url: %s", request.Url)
	body, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("Fetch error: %v, url: %s:", request.Url, err)
		return ParserResult{}, err
	}

	return request.ParserFunc(body), nil
}
