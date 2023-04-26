package test

import (
	"github.com/nsqio/go-nsq"
	"log"
	"testing"
	"time"
)

// 生产者
func startProducer() {
	cfg := nsq.NewConfig()
	producer, err := nsq.NewProducer("1.117.141.66:4150", cfg)
	if err != nil {
		log.Fatal(err)
	}
	// 发布消息
	for {
		if err := producer.Publish("test", []byte("test message")); err != nil {
			log.Fatal("publish error: " + err.Error())
		}
		time.Sleep(3 * time.Second)
	}
}

// 消费者
func startConsumer() {
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("test", "sensor01", cfg)
	if err != nil {
		log.Fatal(err)
	}
	// 设置消息处理函数
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Println(string(message.Body))
		return nil
	}))
	// nsqlookupd
	//[]string
	if err := consumer.ConnectToNSQLookupds([]string{"1.117.141.66:4161"}); err != nil {
		log.Fatal(err)
	}
	<-consumer.StopChan
}

func TestNsq(t *testing.T) {
	startConsumer()
	//startProducer()
}
