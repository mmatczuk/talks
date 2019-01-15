package main

import (
	"fmt"
	"reflect"
)

func main() {
	typeType := reflect.TypeOf((*reflect.Type)(nil)).Elem()
	valueType := reflect.TypeOf((*reflect.Value)(nil))
	fmt.Println(valueType.Implements(typeType))
}
