package main

import (
	"fmt"
	"time"
	"sync/atomic"
)

func printB(i int)  {
fmt.Println("hello world B:",i)
}

var total uint32 =0

func print5(i int) {
	j := i
	for j < i+5 {
		fmt.Println(j)
		j++
	}
	fmt.Println("print5:",i)
	atomic.AddUint32(&total,1)
}

const COUNT =10000

func main() {
	var i  =0
	for ;i<COUNT;i++ {
		go print5(i)
	}

	for i=0;i<COUNT;i++ {
		printB(i)
		//主线程放弃被调度的次数
		time.Sleep(time.Duration(time.Nanosecond))
	}

	fmt.Println("----------total:",total)
}