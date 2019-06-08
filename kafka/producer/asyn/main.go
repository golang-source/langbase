package main


import (
	"github.com/Shopify/sarama"
	"os"
	"log"
	"time"
	"os/signal"
)
//signals := make(chan os.Signal, 1)
//signal.Notify(signals, os.Interrupt)

 var brokerAddrList  = []string {"192.168.1.132:9092"}


var logger = log.New(os.Stderr, "[TEST]", log.LstdFlags)
var normalLog = log.New(os.Stdout, "[Normal]", log.LstdFlags)
var asynProducer sarama.AsyncProducer
var config *sarama.Config

var signals chan os.Signal

func init()  {
	signals = make(chan os.Signal, 1)
	//listen sys signal
	signal.Notify(signals,os.Interrupt)

	config= sarama.NewConfig()
	config.Producer.Return.Successes =false //异步设置为false,如果设置为true，服务器会收不到消息
	config.ClientID = "httpsource"
	//config.Producer.Partitioner = sarama.NewRandomPartitioner //没有设置生产者的分区策略，会导致消息不会发送成功现象
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner //分区策略函数
	asynProducer,_ = newProducer()
}
func  newProducer()  (sarama.AsyncProducer,error)  {
	return sarama.NewAsyncProducer(brokerAddrList,config)
}

var msgCache chan *sarama.ProducerMessage

func createMsg() *sarama.ProducerMessage {
	msgCache = make(chan*sarama.ProducerMessage,100000)

    var msg *sarama.ProducerMessage
for {

	msg = &sarama.ProducerMessage{}
	msg.Topic = "mytopic1"
	msg.Key = sarama.StringEncoder("hello")
	msg.Value = sarama.ByteEncoder("world:http://www.mamicode.com/info-detail-1897291.html")
	msgCache <- msg
	//normalLog.Println("create msg ...")
}
return msg
}
func sendMsg()  {

	for {
		select {
			case msg:=<-msgCache:
			asynProducer.Input() <- msg
		    time.Sleep(100)
				logger.Println("send a msg ...")
		default:
			time.Sleep(100)
		}
	}
}

func listenSinal()  {
	for {
		<-signals
		asynProducer.AsyncClose()
		break
	}
}

func main()  {
	go listenSinal()
	go createMsg()
	go sendMsg()

	select{}

}

