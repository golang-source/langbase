package main

import (
	sarm_cluster "github.com/go-micro-out-parts/sarama-cluster"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
)

func init() {

}
var errLog = log.New(os.Stderr, "[err]", log.LstdFlags)
var normalLog = log.New(os.Stdout, "[Normal]", log.LstdFlags)

var signalChan  chan os.Signal
func listenSignal(consummer *sarm_cluster.Consumer)  {
	signalChan= make(chan os.Signal,1)
	signal.Notify(signalChan,os.Interrupt)
	defer close(signalChan)
	for {
		<-signalChan
		consummer.Close()
	}

}

func pollNofications(consummer *sarm_cluster.Consumer,groupId string)  {
	for noti := range  consummer.Notifications(){
		normalLog.Printf("groupId=%s, noti=%s \n",noti,noti)
	}
}

func pollTopicMsg(consummer *sarm_cluster.Consumer,groupId string)  {
	for {
		select {
				case msg,ok:=  <-consummer.Messages():
					if ok{
						normalLog.Printf("groupId=%s, topic=%s ,partion=%d , key=%s,value=%s , offset=%d \n",groupId,msg.Topic,msg.Partition,msg.Key,msg.Value,msg.Offset)
						consummer.MarkOffset(msg,"")
						}
		}
	}
}

//消费客户端是以组管理的
func newConsummerContext(brokerAddrList,topics []string,groupId string)  {
	 config:=sarm_cluster.NewConfig()
	 //每个消费者
	 config.Consumer.Return.Errors = true
	 config.Consumer.Offsets.Initial = sarama.OffsetNewest
	 //以组的方式有通知
	 config.Group.Return.Notifications = true

	 consummer,err:=sarm_cluster.NewConsumer(brokerAddrList,groupId,topics,config)

	 if err!=nil{
		 errLog.Println(err)
		 return
	 }
	 go listenSignal(consummer)
	 go pollNofications(consummer,groupId)
	 go pollTopicMsg(consummer,groupId)

}


var brokerAddrs =[]string{"192.168.1.132:9092",""}
func main() {
	go newConsummerContext(brokerAddrs,[]string{"mytopic1"},"gorup1")
	go newConsummerContext(brokerAddrs,[]string{"mytopic1"},"gorup2")

	select {
	}
}
