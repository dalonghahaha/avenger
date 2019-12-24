package tests

import (
	"fmt"
	"testing"

	"avenger/tools/uuid"
)

func TestUUID(t *testing.T) {
	fmt.Println(uuid.GenerateV1())
	fmt.Println(uuid.GenerateV4())
	fmt.Println(uuid.Xid())
	fmt.Println(uuid.Sonyflake())
}
