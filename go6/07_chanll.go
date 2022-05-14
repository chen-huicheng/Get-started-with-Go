package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	n := 5
	go func() {
		for i := 0; i < n; i++ {
			fmt.Println("i = ", i)
			ch <- i
		}
		close(ch) //关闭后无法再写数据
	}()
	time.Sleep(time.Second)

	for {
		if num, ok := <-ch; ok == true { //管道关闭 ok 为false
			fmt.Println("num = ", num)
		} else {
			break
		}
	}
	fmt.Println("hello")

	ch1 := make(chan int, 2) //为0会阻塞
	// 双向能隐式转换为单向
	var send chan<- int = ch1 //单向 只写
	var recv <-chan int = ch1 //单向 只读

	send <- 66
	num := <-recv
	// num = <-send //invalid operation: cannot receive from send-only channel send (variable of type chan<- int)
	fmt.Println(num)

	ch2 := make(chan int)
	go producer(ch2)
	consumer(ch2)
}

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}
func consumer(in <-chan int) {
	for num := range in {
		fmt.Println(num)
	}
}
