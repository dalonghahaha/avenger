package coding

import "github.com/axgle/mahonia"

func GBKToUTF8(str string) string {
	dec := mahonia.NewDecoder("gbk")
	return dec.ConvertString(str)
}

func UTF8ToGBK(str string) string {
	enc := mahonia.NewEncoder("gbk")
	return enc.ConvertString(str)
}
