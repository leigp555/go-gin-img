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
}
