package main

import "fmt"

func main() {
	consts()
	enums()
}

func consts() {
	//go语言的首字母大写表示Public，所以常量一般不大写
	const filename = "aaa.txt"
	const a, b = 3, 4
	fmt.Println(filename, a, b)
}

func enums() {
	//在定义常量组时，如果不提供初始值，则表示将使用上行的表达式
	//定义一组const，通过iota实现自增值，这就是常量自增枚举
	const (
		cpp = iota
		_
		python
		golang
		_
		java
	)

	fmt.Println(cpp, python, golang, java)
}
