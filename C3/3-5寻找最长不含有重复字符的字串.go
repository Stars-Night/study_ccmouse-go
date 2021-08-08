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

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		lastI, ok := lastOccurred[ch] //lastOccurred[ch]不存在时取zero value，也就是0，会影响运算。需要判断是否存在
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}
