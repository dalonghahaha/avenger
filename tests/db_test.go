package tests

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"

	"github.com/dalonghahaha/avenger/components/db"
)

func DBInit() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./sample")
	err := viper.ReadInConfig()
	if err != nil {
		panic("go fuck yourself!:" + err.Error())
	}
	err = db.Register()
	if err != nil {
		panic("Register Fail:" + err.Error())
	}
}

func TestDBPing(t *testing.T) {
	DBInit()
	err := db.Get("local").DB().Ping()
	if err != nil {
		t.Error(err)
	}
}

func TestSQLToMap(t *testing.T) {
	DBInit()
	rows, err := db.Get("local").DB().Query("select * from activity_jd limit 1")
	if err != nil {
		t.Error(err)
	}
	data, err := db.ToMap(rows)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*data)
	rows, err = db.Get("local").DB().Query("select * from activity_jd where id = 10086")
	if err != nil {
		t.Error(err)
	}
	data, err = db.ToMap(rows)
	if err != nil && err != db.RowNULL {
		t.Error(err)
	} else if err == db.RowNULL {
		fmt.Println(data)
	}
}

func TestSQLToSliceMap(t *testing.T) {
	DBInit()
	rows, err := db.Get("local").DB().Query("select * from activity_jd limit 5")
	if err != nil {
		t.Error(err)
	}
	datas, err := db.ToSliceMap(rows)
	if err != nil {
		t.Error(err)
	}
	for _, data := range datas {
		fmt.Println(*data)
	}
}
