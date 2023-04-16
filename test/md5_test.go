package test

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func md5Str(str string) (s string) {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

func TestMd5(t *testing.T) {
	s := md5Str("lgp")
	fmt.Println(s)
}
