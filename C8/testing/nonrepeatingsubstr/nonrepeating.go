package main

import "fmt"

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("aaa"))
	fmt.Println(lengthOfNonRepeatingSubStr("abc"))
	fmt.Println(lengthOfNonRepeatingSubStr("abdawet"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("v"))
	fmt.Println(lengthOfNonRepeatingSubStr("你好啊"))
	fmt.Println(lengthOfNonRepeatingSubStr("这z里s是上海!！"))
}

// lastOccurred := make(map[rune]int)
// lastOccurred[ch]	// map access
// lastOccurred[ch] = i	// map assign

func lengthOfNonRepeatingSubStr(s string) int {
	//空间换时间，开辟一个65535个元素，每个元素是int的slice
	//这里假设每个汉字的编码大小最大不会超过0xffff，否则就会超出slice范围报错
	lastOccurred := make([]int, 0xffff)
	for i := range lastOccurred {
		lastOccurred[i] = -1
	}
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) { // []rune(s) is string to slice rune
		lastI := lastOccurred[ch]
		if lastI != -1 && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}
