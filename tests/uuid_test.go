package tests

import (
	"fmt"
	"testing"

	"github.com/dalonghahaha/avenger/tools/uuid"
)

func TestUUID(t *testing.T) {
	fmt.Println(uuid.GenerateV1())
	fmt.Println(uuid.GenerateV4())
	fmt.Println(uuid.Xid())
	fmt.Println(uuid.Sonyflake())
}
