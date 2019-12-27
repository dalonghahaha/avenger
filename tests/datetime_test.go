package tests

import (
	"fmt"
	"testing"

	"github.com/dalonghahaha/avenger/tools/datetime"
)

func TestDateTime(t *testing.T) {
	fmt.Println(datetime.Timestamp())
	fmt.Println(datetime.Now())
	fmt.Println(datetime.Date())
	fmt.Println(datetime.Time())
	fmt.Println(datetime.DateTime())
	fmt.Println(datetime.Parse("May 8, 2009 5:57:51 PM"))
	fmt.Println(datetime.Parse("2006年01月02日"))
}
