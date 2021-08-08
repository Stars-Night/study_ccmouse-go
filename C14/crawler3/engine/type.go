package engine

// Request 请求结构体
type Request struct {
	Url        string                    // 要抓取的url
	ParserFunc func([]byte) ParserResult // 解析抓取内容的方法
}

// ParserResult 解析结果结构体
type ParserResult struct {
	Requests []Request     // 下一步要执行的请求结构体
	Items    []interface{} // 需要的数据
}

// NilParser 一个空的解析方法，调用此方法时，意味着下一步不解析了
func NilParser([]byte) ParserResult {
	return ParserResult{}
}
