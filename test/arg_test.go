package test

import (
	"fmt"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	ch := make(chan int, 10)
	go timeTicker(ch)
	ch <- 100
	time.Sleep(time.Second * 10)
}

func timeTicker(ch chan int) {
	for {
		fmt.Println("准备读取")
		a := <-ch
		fmt.Println(a)
	}
}
