package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"

	"avenger/structs"
)

var dbs = map[string]*gorm.DB{}

var RowNULL = fmt.Errorf("empty result")

func Register() error {
	configs := viper.GetStringMap("component.db")
	for key := range configs {
		config := viper.GetStringMapString("component.db." + key)
		connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config["user"],
			config["password"],
			config["host"],
			config["port"],
			config["database"])
		db, err := gorm.Open("mysql", connStr)
		if err != nil {
			return err
		}
		dbs[key] = db
	}
	return nil
}

func Get(key string) *gorm.DB {
	db, ok := dbs[key]
	if !ok {
		panic("db配置不存在:" + key)
	}
	return db
}

func ToMap(rows *sql.Rows) (*structs.M, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	ok := rows.Next()
	if !ok {
		return nil, RowNULL
	}
	column := make([]interface{}, len(columns))
	args := make([]interface{}, 0)
	for k := range column {
		args = append(args, &column[k])
	}
	err = rows.Scan(args...)
	if err != nil {
		return nil, err
	}
	entry := structs.M{}
	for i, col := range columns {
		var v interface{}
		val := column[i]
		switch val.(type) {
		case []byte:
			v = string(val.([]byte))
		case string:
			v = val
		case time.Time:
			v = val.(time.Time).Format("2006-01-02 15:04:05")
		default:
			if val == nil {
				v = ""
			} else {
				v = val
			}
		}
		entry[col] = v
	}
	return &entry, nil
}

func ToSliceMap(rows *sql.Rows) ([]*structs.M, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	result := []*structs.M{}
	for rows.Next() {
		column := make([]interface{}, len(columns))
		args := make([]interface{}, 0)
		for k := range column {
			args = append(args, &column[k])
		}
		err := rows.Scan(args...)
		if err != nil {
			return nil, err
		}
		entry := structs.M{}
		for i, col := range columns {
			var v interface{}
			val := column[i]
			switch val.(type) {
			case []byte:
				v = string(val.([]byte))
			case string:
				v = val
			case time.Time:
				v = val.(time.Time).Format("2006-01-02 15:04:05")
			default:
				if val == nil {
					v = ""
				} else {
					v = val
				}
			}
			entry[col] = v
		}
		result = append(result, &entry)
	}
	return result, nil
}
