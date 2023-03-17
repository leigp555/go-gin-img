package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5Str(str string) (s string) {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}
