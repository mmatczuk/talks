package main

import "fmt"

func main() {
	var v interface{} // HL
	v = "Hello Gophers"
	v = 42
	fmt.Println(v) // "42"
}
