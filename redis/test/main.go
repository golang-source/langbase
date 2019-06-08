package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

var hosts = []string{"192.168.1.138:2181"}

var path1 = "/ffff"

var flags int32 = zk.FlagSequence
var data1 = []byte("wen weiping zookeeper .....")
var acls = zk.WorldACL(zk.PermAll)


func callback(event zk.Event) {
	fmt.Println("*******************")
	fmt.Println("path:", event.Path)
	fmt.Println("type:", event.Type.String())
	fmt.Println("state:", event.State.String())
	fmt.Println("-------------------")
}

func create(conn *zk.Conn, path string, data []byte) string {
	result, err := conn.Create(path, data, flags, acls)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println("opt:create-result:",result)
	return result
}


func main() {
	option := zk.WithEventCallback(callback)

	conn, _, err := zk.Connect(hosts, time.Second*5, option)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")

	defer conn.Close()

	_, _, _, err = conn.ExistsW(path1)
	if err != nil {
		fmt.Println(err)
		return
	}

	create(conn, path1, data1)


	time.Sleep(time.Second * 2)

	_, _, _, err = conn.ExistsW(path1)
	if err != nil {
		fmt.Println(err)
		return
	}
//	conn.Delete(path1,-1)
	//delete(conn, path1)

fmt.Println("eeeeeeeeeeeeeeeeeeeee")


}

