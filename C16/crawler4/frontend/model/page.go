package model

import "ccmouse-go/C16/crawler4/engine"

type SearchResult struct {
	Hits  int
	Start int
	Items []engine.Item
}
