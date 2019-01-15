package main

import (
	"fmt"
	"reflect"
)

func main() {
	src := []string{"Hello", "Gophers"}
	dst := make([]string, 2)
	reflect.Copy(reflect.ValueOf(dst), reflect.ValueOf(src)) // HL
	fmt.Println(src)
	fmt.Println(dst)
}
