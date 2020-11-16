package kafka_service

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// 基于sarama第三方库开发的kafka client

// Produce is the function of every node to submit their messages to the sync service
func Produce(message []byte) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "quickstart-events"
	msg.Value = sarama.StringEncoder(message)
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return err
	}
	defer client.Close()
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return err
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
	return nil
}
