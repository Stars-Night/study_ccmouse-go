package main

import (
	"fmt"
	"time"
)

func main() {
	demoChannel()
	//bufferedChannel()
	//channelClose()
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	close(c)
	time.Sleep(time.Microsecond)
}

func worker(id int, c <-chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(i int) chan<- int {
	c := make(chan int)
	go worker(i, c)
	return c
}

func demoChannel() {
	var channels [10]chan<- int //定义channel数组
	//开10个goroutine
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Microsecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
}
