package main

import "fmt"

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6} //[]int切片，[...]int是数组
	var arr4 [4][5]int        //二维数组

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(arr4)
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	//数组是值类型
	//[10]int和[20]int是不同类型
	//在go语言中一般不直接使用数组

}
