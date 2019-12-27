package tests

import (
	"fmt"
	"testing"

	"github.com/dalonghahaha/avenger/structs"
	"github.com/dalonghahaha/avenger/tools/http"
)

func TestHttpGet(t *testing.T) {
	data, err := http.New().Get("https://httpbin.org/get")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(data)
}

func TestHttpGetFile(t *testing.T) {
	err := http.New().GetFile("https://cms.avenger.com/images/logo.png", "logo.png")
	if err != nil {
		t.Error(err)
	}
}

func TestHttpPost(t *testing.T) {
	data, err := http.New().PostRaw("https://httpbin.org/post", "hello world")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(data)
}

func TestHttpPostJosn(t *testing.T) {
	data, err := http.New().PostJson("https://httpbin.org/post", structs.M{
		"id":   1,
		"name": "hello",
		"ids":  [...]string{"aa", "bb", "cc"},
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(data)
}

func TestPostForm(t *testing.T) {
	data, err := http.New().PostForm("https://httpbin.org/post", structs.SSM{
		"id":   "hehehhe",
		"name": "hello",
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(data)
}
