package aliyun_oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
)

var clients = map[string]*oss.Client{}

func Register() error {
	configs := viper.GetStringMap("thirdparty.aliyun.oss")
	for key := range configs {
		config := viper.GetStringMapString("thirdparty.aliyun.oss." + key)
		client, err := oss.New(
			config["end_point"],
			config["access_key_id"],
			config["access_key_secret"])
		if err != nil {
			return err
		}
		clients[key] = client
	}
	return nil
}

func Get(key string) *oss.Client {
	client, ok := clients[key]
	if !ok {
		panic("client配置不存在:" + key)
	}
	return client
}
