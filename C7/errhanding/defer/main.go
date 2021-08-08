package main

import (
	"bufio"
	"ccmouse-go/C7/errhanding/defer/fib"
	"fmt"
	"os"
)

func main() {
	writeFile("./fib.txt")
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//bufio会先把内容写到内存里，再一次性写入文件，提升性能
	write := bufio.NewWriter(file)
	defer write.Flush() //写入文件

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(write, f())
	}
}
