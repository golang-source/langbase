package main

import (
	"net/rpc"
	"net"
	"net/http"
	"langbasic/rpc/server"
	"log"
	"fmt"
)

func main() {
	arith := new(servers.Arith) //被注册的struct 必须要有方法
	//public  rpc server class
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":9000")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	fmt.Println("conn:",l)
	go http.Serve(l, nil)

	select {

	}
}