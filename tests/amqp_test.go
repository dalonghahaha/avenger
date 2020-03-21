package tests

import (
	"fmt"
	"log"
	"testing"

	"github.com/spf13/viper"
	_amqp "github.com/streadway/amqp"

	"avenger/components/amqp"
)

func AMQPInit() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./sample/conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic("go fuck yourself!:" + err.Error())
	}
	err = amqp.Register()
	if err != nil {
		panic("Register Fail:" + err.Error())
	}
}

func TestAMQPSendMessage(t *testing.T) {
	AMQPInit()
	conn := amqp.Get("local")
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()
	queue, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Fatal(err)
	}
	body := "Hello World!"
	err = ch.Publish(
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		_amqp.Publishing{ContentType: "text/plain", Body: []byte(body)})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(" [x] Sent %s", body)
}

func TestAMQPReciveMessage(t *testing.T) {
	AMQPInit()
	conn := amqp.Get("local")
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()
	queue, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Fatal(err)
	}
	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		log.Fatal(err)
	}
	for delivery := range msgs {
		fmt.Println("Received a message：", string(delivery.Body))
	}
}
