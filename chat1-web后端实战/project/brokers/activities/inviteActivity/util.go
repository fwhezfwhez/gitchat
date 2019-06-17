package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

// md5 加密
func MD5(rawMsg string) string {
	data := []byte(rawMsg)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	return strings.ToUpper(md5str1)
}
