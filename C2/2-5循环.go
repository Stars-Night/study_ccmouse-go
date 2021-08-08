package main

import "fmt"

func main() {
	for1()
	for2()
	for3()
}

//常规用法
func for1() {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	println(sum)
}

//相当于while，只有结束条件
func for2() {
	i := 10
	for i > 0 {
		fmt.Println(i)
		i--
	}
}

//相当于while 1
func for3() {
	i := 10
	for {
		fmt.Println(i)
		i--
		if i == 5 {
			return
		}
	}
}
