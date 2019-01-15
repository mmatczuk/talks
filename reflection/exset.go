package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 2
	d := reflect.ValueOf(&x).Elem()   // d refers to the variable x
	px := d.Addr().Interface().(*int) // px := &x
	*px = 3                           // x = 3
	fmt.Println(x)

	d.Set(reflect.ValueOf(4))
	fmt.Println(x)

	d.Set(reflect.ValueOf(int64(5)))
}
