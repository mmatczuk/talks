package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.ValueOf(2).Int())
	fmt.Println(reflect.ValueOf(2).Interface())
	fmt.Println(reflect.ValueOf(struct{ f int }{2}).Field(0).Int())
}
