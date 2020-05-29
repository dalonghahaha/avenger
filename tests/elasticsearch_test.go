package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"

	"github.com/dalonghahaha/avenger/components/elasticsearch"
	"github.com/dalonghahaha/avenger/structs"
)

func ESInit() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./sample/conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic("go fuck yourself!:" + err.Error())
	}
	err = elasticsearch.Register()
	if err != nil {
		panic("Register Fail:" + err.Error())
	}
}

func TestESIndexExists(t *testing.T) {
	ESInit()
	client := elasticsearch.Get("local")
	exists, err := client.IndexExists("twitter").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(exists)
}

func TestESIndexList(t *testing.T) {
	ESInit()
	client := elasticsearch.Get("local")
	result, err := client.IndexNames()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestESIndexPut(t *testing.T) {
	ESInit()
	data := `{"user" : "olivere", "message" : "It's a Raggy Waltz"}`
	client := elasticsearch.Get("local")
	put, err := client.Index().Index("test").BodyString(data).Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(put)
}

func TestESIndexQuery(t *testing.T) {
	ESInit()
	client := elasticsearch.Get("local")
	termQuery := elastic.NewTermQuery("user", "olivere")
	search_result, err := client.Search().Index("test").Query(termQuery).From(0).Size(10).Pretty(true).Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if search_result.Hits.TotalHits.Value > 0 {
		result := []structs.M{}
		for _, hit := range search_result.Hits.Hits {
			entry := structs.M{}
			err := json.Unmarshal(hit.Source, &entry)
			if err != nil {
				t.Fatal(err)
			}
			result = append(result, entry)
		}
		fmt.Println(result)
	}
}
