package main

import (
	"github.com/redigo/redis"
	"fmt"
	"os"
)


const REDIS_ADDRESS ="192.168.1.131:6379"
func main() {

	c, err := redis.Dial("tcp", REDIS_ADDRESS)
	if err != nil {
		fmt.Println("connection ")
		os.Exit(-1)
	}


	_,err2:=c.Do("zadd","messagequeue","nx","1.0","buy1","1.1","buy2","1.2","buy3")
	if err2==nil{
		fmt.Println("complete")
	}

	res4,er3:=c.Do("zrange","messagequeue","0","-1")
	strs,ee:=redis.Strings(res4,er3)
	if ee==nil{
		fmt.Println(strs)
	}

	defer c.Close()
}