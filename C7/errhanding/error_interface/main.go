package main

import (
	"bufio"
	"ccmouse-go/C7/errhanding/error_interface/fib"
	"fmt"
	"os"
)

func main() {
	writeFile("fib.txt")
}

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	pathError, ok := err.(*os.PathError)
	if err != nil {
		if !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
		return
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
