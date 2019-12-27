package tests

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/spf13/viper"

	"github.com/dalonghahaha/avenger/components/cache"
	"github.com/dalonghahaha/avenger/tools/coding"
)

func RedisInit() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./sample")
	err := viper.ReadInConfig()
	if err != nil {
		panic("go fuck yourself!:" + err.Error())
	}
	err = cache.Register()
	if err != nil {
		panic("Register Fail:" + err.Error())
	}
}

func TestRedisPing(t *testing.T) {
	RedisInit()
	err := cache.Get("local").Ping().Err()
	if err != nil {
		t.Error(err)
	}
}

func TestRedisCmd(t *testing.T) {
	RedisInit()
	//Set
	err := cache.Get("local").Set("Test", "hello world", time.Hour).Err()
	if err != nil {
		t.Error(err)
	}
	err = cache.Get("local").Set("TestNumber", 10086, time.Hour).Err()
	if err != nil {
		t.Error(err)
	}
	varArr := [...]string{"1", "2", "3"}
	err = cache.Get("local").Set("TestArray", coding.JSONEncode(varArr), time.Hour).Err()
	if err != nil {
		t.Error(err)
	}
	varMap := map[string]interface{}{
		"a": 1,
		"b": 3.1415926,
		"c": "ccc",
		"d": "中华人民共和国",
		"e": "🙄",
	}
	err = cache.Get("local").Set("TestMap", coding.JSONEncode(varMap), time.Hour).Err()
	if err != nil {
		t.Error(err)
	}
	//Get
	result, err := cache.Get("local").Get("Test").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
	result, err = cache.Get("local").Get("TestNumber").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
	result, err = cache.Get("local").Get("TestArray").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
	result, err = cache.Get("local").Get("TestMap").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
	//Exists
	result1, err := cache.Get("local").Exists("Test").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result1)
	result2, err := cache.Get("local").Exists("Test2").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result2)
	//TTL
	result3, err := cache.Get("local").TTL("Test").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result3.Seconds())
	result4, err := cache.Get("local").TTL("Test2").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result4.Seconds())
	//List
	err = cache.Get("local").LPush("TestList", "1", "2", "3", "4", "5", "6", "7", "8").Err()
	if err != nil {
		t.Error(err)
	}
	length, err := cache.Get("local").LLen("TestList").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(length)
	result5, err := cache.Get("local").RPop("TestList").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result5)
	result5, err = cache.Get("local").RPop("TestList").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result5)
	result5, err = cache.Get("local").LPop("TestList").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result5)
	result5, err = cache.Get("local").LPop("TestList").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result5)
	err = cache.Get("local").Del("TestList").Err()
	if err != nil {
		t.Error(err)
	}
	//Hash
	err = cache.Get("local").HSet("TestHash", "aaa", "1111").Err()
	if err != nil {
		t.Error(err)
	}
	err = cache.Get("local").HSet("TestHash", "bbb", "2222").Err()
	if err != nil {
		t.Error(err)
	}
	err = cache.Get("local").HSet("TestHash", "ccc", "3333").Err()
	if err != nil {
		t.Error(err)
	}
	result6, err := cache.Get("local").HGetAll("TestHash").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result6)
	result7, err := cache.Get("local").HGet("TestHash", "aaa").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result7)
	result7, err = cache.Get("local").HGet("TestHash", "bbb").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result7)
	err = cache.Get("local").Del("TestHash").Err()
	if err != nil {
		t.Error(err)
	}
	//Set
	err = cache.Get("local").SAdd("TestSet", "aaa", "bbb", "ccc", "ddd").Err()
	if err != nil {
		t.Error(err)
	}
	result8, err := cache.Get("local").SMembers("TestSet").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result8)
	result9, err := cache.Get("local").SIsMember("TestSet", "aaa").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result9)
	result10, err := cache.Get("local").SCard("TestSet").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result10)
	err = cache.Get("local").SRem("TestSet", "aaa", "bbb").Err()
	if err != nil {
		t.Error(err)
	}
	result8, err = cache.Get("local").SMembers("TestSet").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result8)
	result9, err = cache.Get("local").SIsMember("TestSet", "aaa").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result9)
	result10, err = cache.Get("local").SCard("TestSet").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result10)
	err = cache.Get("local").Del("TestSet").Err()
	if err != nil {
		t.Error(err)
	}
	//ZSet
	err = cache.Get("local").ZAdd("TestZset",
		&redis.Z{Score: 100, Member: "aaa"},
		&redis.Z{Score: 102, Member: "bbb"},
		&redis.Z{Score: 106, Member: "ccc"},
		&redis.Z{Score: 103, Member: "ddd"},
		&redis.Z{Score: 108, Member: "eee"},
		&redis.Z{Score: 104, Member: "fff"},
		&redis.Z{Score: 107, Member: "ggg"},
		&redis.Z{Score: 105, Member: "hhh"},
	).Err()
	if err != nil {
		t.Error(err)
	}
	//获取全部
	result11, err := cache.Get("local").ZRange("TestZset", 0, -1).Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result11)
	//获取数量
	count, err := cache.Get("local").ZCard("TestZset").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(count)
	err = cache.Get("local").ZRem("TestZset", "ggg").Err()
	if err != nil {
		t.Error(err)
	}
	result11, err = cache.Get("local").ZRange("TestZset", 0, -1).Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result11)
	//从小到大取前4
	result11, err = cache.Get("local").ZRange("TestZset", 0, 3).Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result11)
	//从大到小取前4
	result11, err = cache.Get("local").ZRevRange("TestZset", 0, 3).Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result11)
	//获取全部包含分数
	result12, err := cache.Get("local").ZRangeWithScores("TestZset", 0, -1).Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result12)
	//从小到大取前4包含分数
	result12, err = cache.Get("local").ZRangeWithScores("TestZset", 0, 3).Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result12)
	//从大到小取前4包含分数
	result12, err = cache.Get("local").ZRevRangeWithScores("TestZset", 0, 3).Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result12)
	err = cache.Get("local").Del("TestSet").Err()
	if err != nil {
		t.Error(err)
	}
	//原子操作
	var wg sync.WaitGroup
	for index := 0; index < 5000; index++ {
		wg.Add(1)
		go func() {
			cache.Get("local").Incr("TestInc")
			wg.Done()
		}()
	}
	for index := 0; index < 500; index++ {
		wg.Add(1)
		go func() {
			cache.Get("local").Decr("TestInc")
			wg.Done()
		}()
	}
	wg.Wait()
	result13, err := cache.Get("local").Get("TestInc").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result13)
	err = cache.Get("local").Del("TestInc").Err()
	if err != nil {
		t.Error(err)
	}
}
