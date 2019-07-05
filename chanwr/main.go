package main

import (
	"fmt"
	"time"
)

/*func timeoutFun()  {
	//
	sing:=sn.send();
	select {
	   case finish:=  <-sing:
	   	   if finish.aaa{
			   doAck(sing)
		   }

	   	case time.After(time.Second):


	}

}*/

func timeout()  {
	var c =make (chan int)

	go func(cc chan int) {
		time.Sleep(time.Second*5)
		cc<-1
	}(c)

	select {
	   case res:= <-c:
			fmt.Println("res:",res)
	   	case <-time.After(time.Second):
			fmt.Println("timeout")
	}

}


func main() {

	timeout()

	/*
	var a  = make(chan int,0) //size 是给写的缓冲
	go func(c chan int) {
		for i:=0;i<10;i++ {
			c<-i
		}
	}(a)

	for i:=0;i<10;i++ {
		fmt.Println(<-a)
	}*/

}
