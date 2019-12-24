package queue

import (
	"github.com/spf13/viper"
	"gopkg.in/confluentinc/confluent-kafka-go.v100/kafka"
)

var producers = map[string]*kafka.Producer{}
var consumers = map[string]*kafka.Consumer{}

func Register() error {
	producerConfigs := viper.GetStringMap("component.kafka.producer")
	for key := range producerConfigs {
		config := viper.GetStringMapString("component.kafka.producer." + key)
		producer, err := kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": config["broker"],
		})
		if err != nil {
			return err
		}
		producers[key] = producer
	}
	consumerConfigs := viper.GetStringMap("component.kafka.consumer")
	for key := range consumerConfigs {
		config := viper.GetStringMapString("component.kafka.consumer." + key)
		consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": config["broker"],
			"group.id":          "myGroup",
			"auto.offset.reset": "earliest",
		})
		if err != nil {
			return err
		}
		consumers[key] = consumer
	}
	return nil
}

func GetProducer(key string) *kafka.Producer {
	producer, ok := producers[key]
	if !ok {
		panic("producer配置不存在:" + key)
	}
	return producer
}

func GetConsumer(key string) *kafka.Consumer {
	consumer, ok := consumers[key]
	if !ok {
		panic("consumer配置不存在:" + key)
	}
	return consumer
}
