package tests

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"

	"github.com/dalonghahaha/avenger/thirdpartys/aliyun_oss"
)

func OssInit() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./sample/conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic("go fuck yourself!:" + err.Error())
	}
	err = aliyun_oss.Register()
	if err != nil {
		panic("Register Fail:" + err.Error())
	}
}

func TestOssGetBucket(t *testing.T) {
	OssInit()
	result, err := aliyun_oss.Get("qingdao").GetBucketInfo("dalonghahaha")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result.BucketInfo.Name)
	fmt.Println(result.BucketInfo.Location)
	fmt.Println(result.BucketInfo.CreationDate)
	fmt.Println(result.BucketInfo.ExtranetEndpoint)
	fmt.Println(result.BucketInfo.IntranetEndpoint)
	fmt.Println(result.BucketInfo.ACL)
}
