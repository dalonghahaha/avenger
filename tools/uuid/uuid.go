package uuid

import (
	"strconv"

	"github.com/rs/xid"
	UUID "github.com/satori/go.uuid"
	"github.com/sony/sonyflake"
)

//Generate 标准的UUID风格算法 based on timestamp and MAC address (RFC 4122)
func GenerateV1() string {
	return UUID.NewV1().String()
}

//Generate 标准的UUID风格算法 based on random numbers (RFC 4122)
func GenerateV4() string {
	return UUID.NewV4().String()
}

//Xid 基于支持分布式算法：4 bytes time + 3 bytes machine id + 2 bytes process id + 3 bytes random
func Xid() string {
	return xid.New().String()
}

//Sonyflake 基于雪花算法生成
func Sonyflake() string {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		return ""
	}
	return strconv.Itoa(int(id))
}
