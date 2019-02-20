package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	x := os.Stdout
	v := reflect.ValueOf(x)
	fmt.Println(v.Type())
	fmt.Println(v.Elem().Type())
	fmt.Println(v == v.Elem().Addr())
}
