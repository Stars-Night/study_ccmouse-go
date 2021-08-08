package main

import "fmt"

func main() {
	pointer()
}

func pointer() {
	a := 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)

	//go语言指针不能运算，会报错
	//pa++
}
