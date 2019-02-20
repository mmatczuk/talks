package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := reflect.TypeOf(1)
	fmt.Println(i.Size())
	s := reflect.TypeOf("Wrocław loves Go")
	fmt.Println(s.Size())
	e := reflect.TypeOf(struct{}{})
	fmt.Println(e.Size())
	v := reflect.TypeOf([]string{"foo", "bar"})
	fmt.Println(v.Size())
}
