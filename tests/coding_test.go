package tests

import (
	"fmt"
	"testing"

	"github.com/dalonghahaha/avenger/structs"
	"github.com/dalonghahaha/avenger/tools/coding"
)

func TestJSON(t *testing.T) {
	data := structs.M{
		"id":   1,
		"name": "hello world",
		"pi":   3.1415926,
		"isY":  true,
		"info": structs.M{
			"lan": 111,
			"lat": 222,
		},
	}
	fmt.Println(coding.JSONEncode(data))
	jsonStr := `{
		"id":1,
		"name": "hello world",
		"pi":3.1415926,
		"isY":true,
		"info": {
			"lan": 111,
			"lat": 222
		}
	}`
	fmt.Println(coding.JSONDecode(jsonStr))
	jsonStr = `[
		{
			"id":1,
			"name": "juli"
		},
		{
			"id":2,
			"name":"wang"
		}
	]`
	fmt.Println(coding.JSONDecodeArray(jsonStr))
	jsonStr = `{"name":{"first":"Janet","last":"Prichard"},"age":47,"field":{"monney": 3.1415926,"isE": false}}`
	fmt.Println(coding.JSONGet(jsonStr, "name.last"))
	//错误路径(返回对于类型的"零值")
	fmt.Println(coding.JSONGet(jsonStr, "name.test"))
	fmt.Println(coding.JSONGetInt(jsonStr, "age"))
	fmt.Println(coding.JSONGetFloat(jsonStr, "field.monney"))
	fmt.Println(coding.JSONGetBool(jsonStr, "field.isE"))
	fmt.Println(coding.JSONGet(jsonStr, "name"))
	fmt.Println(coding.JSONGetRaw(jsonStr, "field"))
}

func TestXML(t *testing.T) {
	data := structs.M{
		"id":   1,
		"name": "hello world",
		"pi":   3.1415926,
		"isY":  true,
		"info": structs.M{
			"lan": 111,
			"lat": 222,
		},
	}
	fmt.Println(coding.XMLEncode(data))
	xmlStr := `
	<note>
		<to>George</to>
		<from>John</from>
		<heading>Reminder</heading>
	</note>
	`
	fmt.Println(coding.XMLDecode(xmlStr))
	xmlStr = `
	<menu>
		<food>
			<name>Belgian Waffles</name>
			<price>$5.95</price>
			<calories>650</calories>
		</food>
		<food>
			<name>Strawberry Belgian Waffles</name>
			<price>$7.95</price>
			<calories>900</calories>
		</food>
	</menu>
	`
	fmt.Println(coding.XMLDecode(xmlStr))
}

func TestMD5(t *testing.T) {
	str := "hello world"
	coded, err := coding.MD5(str)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(coded)
}

func TestSHA1(t *testing.T) {
	str := "hello world"
	coded, err := coding.SHA1(str)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(coded)
}

func TestSHA2(t *testing.T) {
	str := "hello world"
	coded, err := coding.SHA2(str)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(coded)
}

func TestCRC32(t *testing.T) {
	str := "hello world"
	coded := coding.CRC32(str)
	fmt.Println(coded)
}

func TestCRC64(t *testing.T) {
	str := "hello world"
	coded := coding.CRC64(str)
	fmt.Println(coded)
}

func TestDesEncrypt(t *testing.T) {
	name := "avenger"
	secret := "sdswqeqx"
	key, err := coding.DesEncrypt(name, secret)
	if err == nil {
		fmt.Println(key)
	} else {
		fmt.Println(err)
	}
	//秘钥长度不对测试
	name = "avenger"
	secret = "sds"
	key, err = coding.DesEncrypt(name, secret)
	if err == nil {
		fmt.Println(key)
	} else {
		fmt.Println(err)
	}
}

func TestDesDecrypt(t *testing.T) {
	key := "jNLbvRwV/Mc="
	secret := "sdswqeqx"
	name, err := coding.DesDecrypt(key, secret)
	if err == nil {
		fmt.Println(name)
	} else {
		fmt.Println(err)
	}
	//秘钥长度不对测试
	key = "jNLbvRwV/Mc="
	secret = "sds"
	name, err = coding.DesDecrypt(key, secret)
	if err == nil {
		fmt.Println(name)
	} else {
		fmt.Println(err)
	}
}

func TestAesEncrypt(t *testing.T) {
	name := "avenger"
	secret := "sdswqeqxwedvgthn"
	key, err := coding.AesEncrypt(name, secret)
	if err == nil {
		fmt.Println(key)
	} else {
		fmt.Println(err)
	}
	//秘钥长度不对测试
	name = "avenger"
	secret = "sds"
	key, err = coding.AesEncrypt(name, secret)
	if err == nil {
		fmt.Println(key)
	} else {
		fmt.Println(err)
	}
}

func TestAesDecrypt(t *testing.T) {
	key := "TccuEM3uwXvQcK5E8U1BiA=="
	secret := "sdswqeqxwedvgthn"
	name, err := coding.AesDecrypt(key, secret)
	if err == nil {
		fmt.Println(name)
	} else {
		fmt.Println(err)
	}
	//秘钥长度不对测试
	key = "jNLbvRwV/Mc="
	secret = "sds"
	name, err = coding.AesDecrypt(key, secret)
	if err == nil {
		fmt.Println(name)
	} else {
		fmt.Println(err)
	}
}

func TestConvert(t *testing.T) {
	utfStr := "你好，世界！"
	fmt.Println(coding.UTF8ToGBK(utfStr))
	//"你好，世界！"的GBK编码
	testBytes := []byte{0xC4, 0xE3, 0xBA, 0xC3, 0xA3, 0xAC, 0xCA, 0xC0, 0xBD, 0xE7, 0xA3, 0xA1}
	fmt.Println(coding.GBKToUTF8(string(testBytes)))
}
