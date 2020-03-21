package tests

import (
	"avenger/components/zookeeper"
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

func ZKInit() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./sample/conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic("go fuck yourself!:" + err.Error())
	}
	err = zookeeper.Register()
	if err != nil {
		panic("Register Fail:" + err.Error())
	}
}

func TestZkChildrenW(t *testing.T) {
	ZKInit()
	connect := zookeeper.Get("local")
	children, stat, _, err := connect.ChildrenW("/")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(children)
	fmt.Println(stat)
}
