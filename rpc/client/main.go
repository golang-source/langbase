package main

import "net/rpc"
import "log"
import (
	"fmt"
	"langbasic/rpc/server"
)

func asynCall(client *rpc.Client)  {
	args:=&servers.Args{10,20}
	quotient := new(servers.Quotient)
	divCall:=client.Go("Arith.Divide",args,quotient,nil)
	replyCall:= <-divCall.Done
	fmt.Println("callback",replyCall.Args," ",replyCall.Error)
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, replyCall.Reply)
}

func synCall(client *rpc.Client )  {
	args := &servers.Args{7,8}
	var reply int
	err:= client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
}

const serverAddress="192.168.1.101"
func main() {
	client, err := rpc.DialHTTP("tcp", serverAddress + ":9000")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	//synCall(client)

	asynCall(client)
}

