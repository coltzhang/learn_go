package main

import (
	"fmt"
	"time"
)

var c chan int

func ready(str string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(str, "is ready")
	c <- 1
}

func shower(c chan int, close chan bool) {
	for {
		select {
		case j := <-c:
			fmt.Println(j)

		case <-close:
			break
		}
	}
}

func main() {
	c = make(chan int) // 无缓冲的通道
	go ready("zhangsan", 2)
	go ready("lisi", 1)
	fmt.Println("i am waiting")
	//time.Sleep(time.Duration(5) * time.Second)
	<-c // 阻塞等待，直到有数据可读
	_, ok := <-c
	fmt.Println(ok)

	close := make(chan bool)
	go shower(c, close)
	for i := 0; i < 10; i++ {
		c <- i
		time.Sleep(1 * time.Millisecond)
	}
	close <- true
}
