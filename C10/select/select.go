package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, c <-chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func main() {
	var c1, c2 = generator(), generator() //生产数据的channel
	worker := createWorker(0)             //消费数据的channel

	n := 0
	var activeWorker chan<- int
	var values []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {
		activeWorker = nil //chan变量为nil时无法通信
		var activeValue int
		if len(values) > 0 {
			activeValue = values[0]
			activeWorker = worker //n有数据时activeWorker才发送数据
		}

		select {
		case n = <-c1:
			values = append(values, n)
		case n = <-c2:
			values = append(values, n)
		case activeWorker <- activeValue: //单独开一个case保证发数据不阻塞
			values = values[1:]
		case <-time.After(800 * time.Millisecond): //800毫秒其它case不执行就会打印timeout
			fmt.Println("timeout")
		case <-tick: //每秒计算一次values的长度
			fmt.Println("values len = ", len(values))
		case <-tm: //10秒后退出
			fmt.Println("bye")
			return
		}
	}
}
