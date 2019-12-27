package tests

import (
	"fmt"
	"testing"

	"github.com/dalonghahaha/avenger/tools/random"
)

func TestRandom(t *testing.T) {
	fmt.Println(random.Rand(100))
	fmt.Println(random.Between(100, 150))
	fmt.Println(random.Chars(100))
	fmt.Println(random.CharsPlus(100))
	fmt.Println(random.Numbers(100))
	fmt.Println(random.Letters(100))
	fmt.Println(random.LowerCases(100))
	fmt.Println(random.UpperCases(100))
}
