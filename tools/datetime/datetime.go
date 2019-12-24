package datetime

import (
	"time"

	"github.com/araddon/dateparse"
)

//Timestamp 获取当前时间戳
func Timestamp() int64 {
	return Now().Unix()
}

//Now 获取当前时间(默认北京时间)
func Now() time.Time {
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		location = time.Local
	}
	return time.Now().In(location)
}

//DateTime 获取当前时间(格式：2006-01-02)
func DateTime() string {
	return Now().Format("2006-01-02 15:04:05")
}

//Date 获取当前时间日期(格式：2006-01-02)
func Date() string {
	return Now().Format("2006-01-02")
}

//Time 获取当前时间时间(格式：15:04:05)
func Time() string {
	return Now().Format("15:04:05")
}

//Parse 时间格式格式解析函数(兼容大部分时间字符串格式，对中文不太友好)
func Parse(str string) (time.Time, error) {
	return dateparse.ParseAny(str)
}
