package test

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

var (
	arr = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
)

func TestNackName(t *testing.T) {
	rand.Seed(time.Now().Unix())
	nackArr := make([]string, 5)
	for i := 0; i < 5; i++ {
		nackArr = append(nackArr, arr[rand.Intn(len(arr))])
	}
	fmt.Println(strings.Join(nackArr, ""))
}
