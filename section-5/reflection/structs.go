package main

import (
	"fmt"
	"reflect"
)

func main() {
	type myStruct struct {
		Field1 int     `alias:"f1" desc:"field number 1"`
		Field2 string  `alias:"f2" desc:"field number 2"`
		Field3 float64 `alias:"f3" desc:"field number 3"`
	}
	mys := myStruct{2, "Hello", 2.4}

	mysRValue := reflect.ValueOf(mys)
	mysRType := reflect.TypeOf(mys)

	for i := 0; i < mysRType.NumField(); i++ {
		fieldRType := mysRType.Field(i)
		fieldRValue := mysRValue.Field(i)

		fmt.Printf(
			"Field Name: '%s', field type: '%s', field value: '%v', \n",
			fieldRType.Name,
			fieldRType.Type,
			fieldRValue.Interface(),
		)
		fmt.Println(
			"Struct tags, alias: ",
			fieldRType.Tag.Get("alias"),
			" description: ",
			fieldRType.Tag.Get("desc"),
		)
	}
}
