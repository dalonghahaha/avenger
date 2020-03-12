package elasticsearch

import (
	"fmt"

	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
)

var clients = map[string]*elastic.Client{}

func Register() error {
	configs := viper.GetStringMap("component.elasticsearch")
	for key := range configs {
		config := viper.GetStringMapString("component.elasticsearch." + key)
		connStr := fmt.Sprintf("http://%s:%s", config["host"], config["port"])
		option := []elastic.ClientOptionFunc{}
		option = append(option, elastic.SetSniff(false))
		option = append(option, elastic.SetHealthcheck(false))
		option = append(option, elastic.SetURL(connStr))
		client, err := elastic.NewClient(option...)
		if err != nil {
			return err
		}
		clients[key] = client
	}
	return nil
}

func Get(key string) *elastic.Client {
	client, ok := clients[key]
	if !ok {
		panic("elasticsearch配置不存在:" + key)
	}
	return client
}
