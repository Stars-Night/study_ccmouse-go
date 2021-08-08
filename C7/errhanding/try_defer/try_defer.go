package main

import "fmt"

func main() {
	tryDefer()
}

//defer 栈
//defer能保证程序在不正常return或panic时也能执行某些操作，比如关闭资源
func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occurred")
	return
	fmt.Println(4)
}
