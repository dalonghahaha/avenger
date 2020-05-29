package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"gopkg.in/mgo.v2/bson"

	"github.com/dalonghahaha/avenger/components/mongo"
)

func MongoInit() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./sample/conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic("go fuck yourself!:" + err.Error())
	}
	err = mongo.Register()
	if err != nil {
		panic("Register Fail:" + err.Error())
	}
}

func TestMongoListDatabaseNames(t *testing.T) {
	MongoInit()
	client := mongo.Get("local")
	result, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		t.Fatal(err)
	}
	for _, db := range result {
		fmt.Println(db)
	}
}
