package main

import (
	"fmt"
	"time"
	"sync"
)


var total uint32=0
var rwLock sync.RWMutex
func plusOpt()  {
	rwLock.Lock()
	total++
	rwLock.Unlock()
}

func newGoroutine(j int)  {
		for i:=0;i<100;i++ {
			fmt.Println("hello world:gid:",j," taskid:",i)
			plusOpt()
			time.Sleep(time.Duration(time.Nanosecond))
		}
	wg.Done()
}



var wg sync.WaitGroup

func main() {
	rwLock = sync.RWMutex{}
	wg = sync.WaitGroup{}
	wg.Add(100)
	for i:=0;i<100;i++{
		go newGoroutine(i)
	}
	wg.Wait()
    fmt.Println("total :",total)

}
