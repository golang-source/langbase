package main

import (
	"time"
	"fmt"
)

func onRender()  {
	fmt.Println("on render msg")
}

func main() {
	ticker:=time.NewTicker(time.Duration(time.Second))

	var c =0

	for {
		select {
		case t := <-ticker.C:
			c++
			onRender()

			fmt.Println(t)
			if c == 10 {
				goto endflag
			}
		}
	}
		endflag:
			fmt.Println("complete...")
			ticker.Stop()

}
