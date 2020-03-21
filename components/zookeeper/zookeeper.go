package zookeeper

import (
	"strings"
	"time"

	"github.com/samuel/go-zookeeper/zk"
	"github.com/spf13/viper"
)

var connects = map[string]*zk.Conn{}

func Register() error {
	configs := viper.GetStringMap("component.zookeeper")
	for key := range configs {
		config := viper.GetStringMapString("component.zookeeper." + key)
		hosts := strings.Split(config["hosts"], ";")
		connect, _, err := zk.Connect(hosts, time.Second, zk.WithLogInfo(false))
		if err != nil {
			return err
		}
		connects[key] = connect
	}
	return nil
}

func Get(key string) *zk.Conn {
	connect, ok := connects[key]
	if !ok {
		panic("zookeeper配置不存在:" + key)
	}
	return connect
}
