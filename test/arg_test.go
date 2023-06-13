package test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect(x *testing.T) {
	a := 1
	//t:=reflect.TypeOf(a)
	v := reflect.ValueOf(&a)
	v.Elem().SetInt(111)
	fmt.Println(a)

	b := 10
	if b > 2 {
		fmt.Println("b>2")
	} else if b > 5 {
		fmt.Println("b>5")
	} else if b > 100 {
		fmt.Println("b>100")
	}
}
