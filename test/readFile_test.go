package test

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestReader(t *testing.T) {
	r, err := os.OpenFile("../settings.yaml", os.O_RDWR, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	newR := bufio.NewReader(r)
	buf := make([]byte, 10)
	for {
		n, err := newR.Read(buf) //读到缓冲区
		if err == io.EOF {
			break
		} else {
			fmt.Println(string(buf[0:n]))
		}
	}

}
