package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱你中国!" //UTF-8
	fmt.Println(s)

	//把s转换成byte类型的slice
	for _, ch := range []byte(s) { //ch是utf-8解码后的编码，一个汉字三个字节
		fmt.Printf("%X ", ch)
	}
	fmt.Println()

	// 字符串直接range得到的值是rune类型，即ch的类型是rune
	// 但i的值对应的是字节的起始位置，而不是字符
	for i, ch := range s { //ch的内容是utf-8解码后转成的unicode编码，一个汉字2个字节
		fmt.Printf("(%d %c)", i, ch)
		//fmt.Println(reflect.TypeOf(ch))
	}
	fmt.Println()

	fmt.Println(utf8.RuneCountInString(s)) //获得utf8的字符数，len只能获得字节数

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes) //ch是rune类型的Unicode码点
		//fmt.Println(reflect.TypeOf(ch))
		bytes = bytes[size:]  //reslice
		fmt.Printf("%c ", ch) //%c表示Unicode码点对应的字符
	}
	fmt.Println()

	//遍历中文字符串需要把字符串转换成rune类型的slice，这样i, ch就是一一对应
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()

}
