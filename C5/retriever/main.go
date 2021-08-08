package main

import (
	"ccmouse-go/C5/retriever/mooc"
	"fmt"
)

const url = "http://www.imooc.com"

type Retriever interface {
	Get(url string) string
}

func Download(r Retriever) string {
	return r.Get(url)
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func Post(poster Poster) string {
	return poster.Post(url, map[string]string{
		"name":     "zhangsan",
		"contents": "lalala",
	})
}

//组合接口
type RetrieverPoster interface {
	Retriever //接口
	Poster    //接口
	//其它方法
	//Request() string
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)

}

func main() {
	var r Poster                                                //定义一个Poster接口类型变量
	retrieverPoster := mooc.Retriever{"this is fake imooc.com"} //把mooc.Retriever对象赋给RetrieverPoster接口类型变量
	fmt.Println(&retrieverPoster)
	fmt.Println(retrieverPoster.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	}))
	fmt.Println(Download(&retrieverPoster))

	r = &retrieverPoster
	fmt.Printf("%T %v\n", r, r)

	fmt.Println("try a session")
	fmt.Println(session(&retrieverPoster))
}
