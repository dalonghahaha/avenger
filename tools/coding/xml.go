package coding

import (
	"github.com/dalonghahaha/avenger/structs"

	"github.com/clbanning/mxj"
)

//XMLEncode Xml格式编码
func XMLEncode(value interface{}) string {
	info := JSONEncode(value)
	v := JSONDecode(info)
	bytes, err := mxj.Map(v).Xml()
	if err != nil {
		return ""
	}
	return string(bytes)
}

//XMLDecode Xml格式解析
func XMLDecode(str string) structs.M {
	m := structs.M{}
	info, err := mxj.NewMapXml([]byte(str))
	if err != nil {
		return structs.M{}
	}
	for k, v := range info {
		m[k] = v
	}
	return m
}