package coding

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"hash/crc32"
	"hash/crc64"
)

func MD5(str string) (string, error) {
	bytes, err := hashBytes("md5", str)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func SHA1(str string) (string, error) {
	bytes, err := hashBytes("sha1", str)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func SHA2(str string) (string, error) {
	bytes, err := hashBytes("sha2", str)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func CRC32(str string) int {
	tab := crc32.MakeTable(crc32.IEEE)
	coded := crc32.Checksum([]byte(str), tab)
	return int(coded)
}

func CRC64(str string) int64 {
	tab := crc64.MakeTable(crc64.ISO)
	coded := crc64.Checksum([]byte(str), tab)
	return int64(coded)
}

func hashBytes(algorithm, str string) ([]byte, error) {
	var h hash.Hash
	switch algorithm {
	case "md5":
		h = md5.New()
	case "sha1":
		h = sha1.New()
	case "sha2", "sha256":
		h = sha256.New()
	}
	_, err := h.Write([]byte(str))
	if err != nil {
		return []byte{}, err
	}
	return h.Sum(nil), nil
}
