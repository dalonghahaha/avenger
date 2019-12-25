package aliyun_ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/spf13/viper"
)

var clients = map[string]*ecs.Client{}

func Register() error {
	configs := viper.GetStringMap("thirdparty.aliyun.ecs")
	for key := range configs {
		config := viper.GetStringMapString("thirdparty.aliyun.ecs." + key)
		client, err := ecs.NewClientWithAccessKey(
			config["region_id"], 
			config["access_key_id"], 
			config["access_key_secret"])
		if err != nil {
			return err
		}
		clients[key] = client
	}
	return nil
}

func Get(key string) *ecs.Client {
	client, ok := clients[key]
	if !ok {
		panic("client配置不存在:" + key)
	}
	return client
}
