package engine

import (
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

		parserResult, err := Worker(request)
		if err != nil {
			continue
		}

		requests = append(requests, parserResult.Requests...)
		for _, item := range parserResult.Items {
			log.Printf("got item: %v", item)
		}
	}
}
