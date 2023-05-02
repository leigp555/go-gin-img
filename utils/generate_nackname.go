package utils

import (
	"math/rand"
	"strings"
	"time"
)

var (
	arr = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
)

func GenerateNackName() string {
	rand.Seed(time.Now().Unix())
	nackArr := make([]string, 5)
	for i := 0; i < 5; i++ {
		nackArr = append(nackArr, arr[rand.Intn(len(arr))])
	}
	return strings.Join(nackArr, "")
}
