package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 2
	fmt.Println(reflect.ValueOf(2).CanAddr())
	fmt.Println(reflect.ValueOf(x).CanAddr())
	fmt.Println(reflect.ValueOf(&x).CanAddr())
	fmt.Println(reflect.ValueOf(&x).Elem().CanAddr())
}
