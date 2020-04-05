package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x1 float32 = 5.6

	v := reflect.ValueOf(x1)
	fmt.Println("type: ", v.Type())
	fmt.Println("Is type float64?", v.Kind() == reflect.Float64)
	fmt.Println("Float Value:", v.Float())
}