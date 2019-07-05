package main

import (
	"runtime"
	"fmt"
	"sync/atomic"
)

const NUM =10
const FORSIZE= 10000

func routine(t *int32,clock chan int)  {

	for i:=0;i<FORSIZE;i++ {
		//(*t)++			//并发环境+非原子操作 导致
		atomic.AddInt32(t,1)
		runtime.Gosched()
	}
	clock<-1
}




func main()  {
	var i int
	var total int32 =0
	var recvWaterTrigger =make(chan int,1) //缓冲大小，对写端有卡主 堵塞的作用，对于读端，没东西就卡住，堵塞
	//所以 channel 对读写都有 堵塞作用的
	for i=0;i<NUM;i++{
		go routine(&total,recvWaterTrigger)
	}

	for i=0;i<NUM;i++{
			<-recvWaterTrigger	//间歇性 堵塞，接触 有 对端控制
	}
	fmt.Printf("total=%d \n",total)
	
}
