package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := reflect.TypeOf(1)
	fmt.Println(i.Size())
	s := reflect.TypeOf("Golang Warsaw")
	fmt.Println(s.Size())
	e := reflect.TypeOf(struct{}{})
	fmt.Println(e.Size())
}
