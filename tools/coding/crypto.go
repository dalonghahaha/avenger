package coding

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"errors"
)

func pkcsPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcsUnpadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//Base64Encode Base64格式编码
func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

//Base64Decode Base64格式解码
func Base64Decode(src string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

//DesEncrypt DES算法加密
func DesEncrypt(str, key string) (string, error) {
	_str := []byte(str)
	_key := []byte(key)
	if len(_key) != 8 {
		return "", errors.New("key size must 8")
	}
	block, err := des.NewCipher(_key)
	if err != nil {
		return "", err
	}
	origData := pkcsPadding(_str, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, _key)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return Base64Encode(string(crypted)), nil
}

//DesDecrypt DES算法解密
func DesDecrypt(str, key string) (string, error) {
	_key := []byte(key)
	if len(_key) != 8 {
		return "", errors.New("key size must 8")
	}
	crypted, err := Base64Decode(str)
	if err != nil {
		return "", errors.New("base64 decode fail")
	}
	if len(crypted)%des.BlockSize != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	block, err := des.NewCipher(_key)
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, _key)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, []byte(crypted))
	origData = pkcsUnpadding(origData)
	return string(origData), nil
}

//AesEncrypt AES算法加密
func AesEncrypt(str, key string) (string, error) {
	_str := []byte(str)
	_key := []byte(key)
	if len(_key) != 16 {
		return "", errors.New("key size must 16")
	}
	block, err := aes.NewCipher(_key)
	if err != nil {
		return "", err
	}
	origData := pkcsPadding(_str, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, _key)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return Base64Encode(string(crypted)), nil
}

//AesDecrypt AES算法解密
func AesDecrypt(encodeString string, key string) (string, error) {
	_key := []byte(key)
	if len(_key) != 16 {
		return "", errors.New("key size must 16")
	}
	crypted, err := Base64Decode(encodeString)
	if err != nil {
		return "", errors.New("base64 decode fail")
	}
	if len(crypted)%aes.BlockSize != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	block, err := aes.NewCipher(_key)
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, _key)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, []byte(crypted))
	origData = pkcsUnpadding(origData)
	return string(origData), nil
}
