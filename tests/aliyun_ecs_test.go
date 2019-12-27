package tests

import (
	"fmt"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/spf13/viper"

	"github.com/dalonghahaha/avenger/thirdpartys/aliyun_ecs"
)

func EcsInit() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./sample/conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic("go fuck yourself!:" + err.Error())
	}
	err = aliyun_ecs.Register()
	if err != nil {
		panic("Register Fail:" + err.Error())
	}
}

func TestEcsDescribeInstance(t *testing.T) {
	EcsInit()
	request := ecs.CreateDescribeInstanceAttributeRequest()
	request.Scheme = "https"
	request.InstanceId = "AY140328105720046324"
	result, err := aliyun_ecs.Get("qingdao").DescribeInstanceAttribute(request)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result.InstanceId)
	fmt.Println(result.InstanceName)
}

func TestEcsDescribeDisks(t *testing.T) {
	EcsInit()
	request := ecs.CreateDescribeDisksRequest()
	request.Scheme = "https"
	result, err := aliyun_ecs.Get("qingdao").DescribeDisks(request)
	if err != nil {
		t.Error(err)
	}
	for _, v := range result.Disks.Disk {
		fmt.Println()
		fmt.Println(v.DiskId)
		fmt.Println(v.DiskName)
		fmt.Println(v.Size)
	}
}
