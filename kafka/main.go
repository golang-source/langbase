package main

import (
	"github.com/Shopify/sarama"
	"log"
	"os"
	"strings"

)

var logger = log.New(os.Stderr, "[TEST]", log.LstdFlags)
var normalLog = log.New(os.Stdout, "[Normal]", log.LstdFlags)



func main(){
	sarama.Logger = logger

	config := sarama.NewConfig()
	config.ClientID = "newsDataSource"
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner //指定主题分区策略函数

	msg := &sarama.ProducerMessage{}
	msg.Topic = "mytopic1"
	msg.Partition = int32(-1)
	msg.Key = sarama.StringEncoder("key")
	msg.Value = sarama.ByteEncoder("hello")

	producer, err := sarama.NewSyncProducer(strings.Split("192.168.1.132:9092", ","), config)
	if err != nil {
		logger.Printf("Failed to produce message :%s", err )
		os.Exit(500)
	}

	defer producer.Close()

	for i:=0;i<100;i++ {
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			logger.Printf("Failed to produce message :%s", err )
		}

		normalLog.Printf("partition:%d, offset: %d\n", partition, offset )
		//time.Sleep(time.Duration(1000*time2.Millisecond))
	}

}