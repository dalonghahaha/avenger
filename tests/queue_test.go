package tests

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"gopkg.in/confluentinc/confluent-kafka-go.v100/kafka"

	"github.com/dalonghahaha/avenger/components/queue"
)

func QueueInit() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./sample/conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic("go fuck yourself!:" + err.Error())
	}
	err = queue.Register()
	if err != nil {
		panic("Register Fail:" + err.Error())
	}
}

func TestQueueSendMessage(t *testing.T) {
	QueueInit()
	producer := queue.GetProducer("local")
	topic := "myTopic"
	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: []byte(word),
		}, nil)
	}
	producer.Flush(15 * 1000)
}

func TestQueueReciveMessage(t *testing.T) {
	QueueInit()
	consumer := queue.GetConsumer("local")
	consumer.SubscribeTopics([]string{"myTopic"}, nil)
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
