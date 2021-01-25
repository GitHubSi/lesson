package main

import (
	"fmt"
	"github.com/spf13/cast"
	"reflect"
)

func main() {
	//cmd.Execute()


	type Infos []string

	personA := Infos{"fuhui"}
	paramsKv := reflect.ValueOf(personA)

	paramIndexV := paramsKv.Index(0)
	fmt.Println(paramIndexV)

	paramIndexVal := reflect.ValueOf(paramIndexV.Interface())
	name := cast.ToString(paramIndexVal.Interface())
	fmt.Println(name)

	name2 := cast.ToString(paramIndexV.Interface())
	fmt.Println(name2)
}
