package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := []string{"Hello", "Gophers"}
	fmt.Println(reflect.TypeOf(m).Elem())
	fmt.Println(reflect.ValueOf(m).Elem())
}
