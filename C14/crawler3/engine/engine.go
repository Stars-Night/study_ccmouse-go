package engine

import (
	"ccmouse-go/C14/crawler3/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		log.Printf("Fetching Url: %s", request.Url)
		body, err := fetcher.Fetch(request.Url)
		if err != nil {
			log.Printf("Fetch error: %v, url: %s:", request.Url, err)
			continue
		}

		parserResult := request.ParserFunc(body)
		requests = append(requests, parserResult.Requests...)
		for _, item := range parserResult.Items {
			log.Printf("got item: %v", item)
		}
	}
}
