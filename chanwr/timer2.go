package main

import (
	"fmt"

	"time"
)

func main() {

	t := time.NewTimer(time.Second * 2)
	defer t.Stop()
	var i=0
	for {
		<-t.C
		fmt.Println("timer running...")
		// 需要重置Reset 使 t 重新开始计时
		t.Reset(time.Second * 2)
		i++
		if i==10{
			break
		}
	}

	//如果主线程 一直得不到，调度权，就会死锁
}
