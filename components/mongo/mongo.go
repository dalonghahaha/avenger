package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var clients = map[string]*mongo.Client{}

func Register() error {
	configs := viper.GetStringMap("component.mongo")
	for key := range configs {
		config := viper.GetStringMapString("component.mongo." + key)
		connStr := fmt.Sprintf("mongodb://%s:%s", config["host"], config["port"])
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(connStr))
		if err != nil {
			return err
		}
		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			return err
		}
		clients[key] = client
	}
	return nil
}

func Get(key string) *mongo.Client {
	client, ok := clients[key]
	if !ok {
		panic("mongo配置不存在:" + key)
	}
	return client
}
