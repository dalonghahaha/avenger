package amqp

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

var connections = map[string]*amqp.Connection{}

func Register() error {
	configs := viper.GetStringMap("component.amqp")
	for key := range configs {
		config := viper.GetStringMapString("component.amqp." + key)
		conStr := fmt.Sprintf("amqp://%s:%s@%s:%s/",
			config["user"],
			config["password"],
			config["host"],
			config["port"],
		)
		conn, err := amqp.Dial(conStr)
		if err != nil {
			return err
		}
		connections[key] = conn
	}
	return nil
}

func Get(key string) *amqp.Connection {
	connection, ok := connections[key]
	if !ok {
		panic("amqp配置不存在:" + key)
	}
	return connection
}
