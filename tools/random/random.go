package random

import (
	"math/rand"
	"time"
)

var (
	digits     = []rune("0123456789")
	lowerCases = []rune("abcdefghijklmnopqrstuvwxyz")
	upperCases = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	englishs   = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	letters    = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	letterPlus = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*~|")
)

//Rand 生成随机数，大于0小于max
func Rand(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max)
}

//Between 生成随机数，在min和max之间
func Between(min int, max int) int {
	return min + Rand(max-min)
}

//String 生成随机n位字符串(包含数字或者字符)
func Chars(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[Rand(62)]
	}
	return string(b)
}

//String 生成随机n位字符串(包含数字或者字符和特殊字符)
func CharsPlus(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterPlus[Rand(72)]
	}
	return string(b)
}

//String 生成随机n位字符串(只包含数字)
func Numbers(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = digits[Rand(10)]
	}
	return string(b)
}

//String 生成随机n位字符串(只包含字符)
func Letters(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = englishs[Rand(52)]
	}
	return string(b)
}

//String 生成随机n位字符串(只包含小写字符)
func LowerCases(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = lowerCases[Rand(26)]
	}
	return string(b)
}

//String 生成随机n位字符串(只包含小写字符)
func UpperCases(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = upperCases[Rand(26)]
	}
	return string(b)
}
