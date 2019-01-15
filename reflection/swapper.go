package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	f := reflect.Swapper(s) // HL
	fmt.Println(s)
	f(0, 5)
	fmt.Println(s)
}
