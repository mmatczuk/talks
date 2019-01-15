package main

import (
	"fmt"
	"reflect"
)

func main() {
	type A struct {
		V int
		B struct {
			V int
			v int
		}
	}

	k := A{}
	k.V = 1
	k.B.V = 2
	k.B.v = 3 // HL

	l := A{}
	l.V = 1
	l.B.V = 2
	l.B.v = 4 // HL

	fmt.Println(reflect.DeepEqual(k, l))
}
