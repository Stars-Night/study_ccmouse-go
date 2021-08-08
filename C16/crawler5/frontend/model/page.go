package model

//定义页面所需的数据结构
type SearchResult struct {
	Hits  int64
	Start int
	Items []interface{}
}
