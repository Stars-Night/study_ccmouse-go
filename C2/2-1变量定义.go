package main

import "fmt"

//函数体外部只能声明语句（即语句前面必须带关键字）
var aa = 11

//否则会报错 syntax error: non-declaration statement outside function body（函数体外部的非声明语句）
//错误示例：
//fmt.Println("hello")
//aa = bb + cc

func main() {
	variableZeroValue()
	variableInitialValue()
	variableShorter()
}

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b = 3, 4
	var s = "abc"
	fmt.Println(a, b, s)
}

//一般定义变量都用这种
func variableShorter() {
	a, b := 3, 4
	s := "abc"
	t := true
	fmt.Println(a, b, s, t)
}
