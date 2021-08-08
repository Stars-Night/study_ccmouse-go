package main

import (
	"ccmouse-go/C16/crawler5/frontend/controller"
	"net/http"
)

func main() {
	//映射网站根目录的物理地址，以便访问静态资源：css、js等
	http.Handle("/", http.FileServer(http.Dir("C16/crawler5/frontend/view")))
	//起一个http服务，浏览器可访问
	http.Handle("/search", controller.CreateSearchResultHandler("C16/crawler5/frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
