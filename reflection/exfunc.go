package main

import (
	"fmt"
	"reflect"
)

func main() {
	f := func(a, b int) int { return a + b }
	funType := reflect.TypeOf(f)
	fmt.Println(funType.NumIn())
	fmt.Println(funType.In(0))
}
