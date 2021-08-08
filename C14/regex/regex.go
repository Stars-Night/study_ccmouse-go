package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is ccmouse@gmail.com@ew12.com
email1 is wwwww@qq.com
email2 is dsrffew@163.com
email3 is dsrffew@asd3.com.cn
`

func main() {
	// ()：标记一个子表达式的开始和结束位置。
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	//match := re.FindString(text)	//只返回第一个匹配的字符
	match := re.FindAllStringSubmatch(text, -1) //返回所有匹配的字符	Submatch：子匹配
	fmt.Println(match)
}
