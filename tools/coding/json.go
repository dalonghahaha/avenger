package coding

import (
	"bytes"
	"encoding/json"

	"github.com/tidwall/gjson"

	"avenger/structs"
)

//JSONEncode Json编码
func JSONEncode(value interface{}) string {
	_bytes, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(_bytes)
}

//JSONDecode Json解析(返回structs.M)
func JSONDecode(str string) structs.M {
	result := structs.M{}
	if len(str) <= 0 {
		return result
	}
	d := json.NewDecoder(bytes.NewReader([]byte(str)))
	d.UseNumber()
	err := d.Decode(&result)
	if err != nil {
		return structs.M{}
	}
	return result
}

//JSONDecodeArray Json数组解析(返回[]structs.M)
func JSONDecodeArray(str string) []structs.M {
	result := []structs.M{}
	if len(str) <= 0 {
		return result
	}
	d := json.NewDecoder(bytes.NewReader([]byte(str)))
	d.UseNumber()
	err := d.Decode(&result)
	if err != nil {
		return []structs.M{}
	}
	return result
}

//JSONGet 快捷方式访问通过路径访问属性值(返回interface{})
func JSONGet(json string, path string) interface{} {
	return gjson.Get(json, path).Value()
}

//JSONGetString 快捷方式访问通过路径访问属性值(返回字符串)
func JSONGetString(json string, path string) string {
	return gjson.Get(json, path).String()
}

//JSONGetInt 快捷方式访问通过路径访问属性值(返回整数)
func JSONGetInt(json string, path string) int64 {
	return gjson.Get(json, path).Int()
}

//JSONGetInt 快捷方式访问通过路径访问属性值(返回浮点数)
func JSONGetFloat(json string, path string) float64 {
	return gjson.Get(json, path).Float()
}

//JSONGetBool 快捷方式访问通过路径访问属性值(返回bool值)
func JSONGetBool(json string, path string) bool {
	return gjson.Get(json, path).Bool()
}

//JSONGetRaw 快捷方式访问通过路径访问属性值(返回字符串,适合属性仍然是Json对象的场景)
func JSONGetRaw(json string, path string) string {
	return gjson.Get(json, path).Raw
}
