package main

import (
	"fmt"
	"time"
)

func addNumberToChan(chanName chan int) {
	for {
		chanName <- 1
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var chan1 = make(chan int, 10)
	var chan2 = make(chan int, 10)

	go addNumberToChan(chan1)
	go addNumberToChan(chan2)

	//如果select 没有default ,select{}检测 ，没有侦听到消息，就会把当前线程加入到一个 waitq
	for {
		select {
		case e := <- chan1 ://当前线程如果没有获得缓冲channel的数据，那么当前线程会被加入到一个对方的 读等待队列，这里是 chan1的读等待队列
			fmt.Printf("Get element from chan1: %d\n", e)
		case e := <- chan2 ://这里是 chan2的读等待队列
			fmt.Printf("Get element from chan2: %d\n", e)
		default://如果加了 default 那就会影响当前的策略，有default，如果检测 select{上下} 无数据可以读，就会直接返回，不会当前线程不会加入到对方的等待队列
			fmt.Printf("No element in chan1 and chan2.\n")
			time.Sleep(1 * time.Second)
		}
	}
}