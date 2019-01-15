package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	stdout := reflect.ValueOf(os.Stdout).Elem() // *os.Stdout, an os.File var
	fmt.Println(stdout.Type())                  // "os.File"
	fd := stdout.Field(0).Elem().FieldByName("pfd").FieldByName("Sysfd")
	fmt.Println(fd.CanAddr())
	fmt.Println(fd.Int())
	fd.SetInt(2) // HL
}
