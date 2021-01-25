package main

import (
	"fmt"
	"math"
)

type inner struct {
	Msg string
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Print(r.(inner))
		}
	}()

	// panic(inner{Msg: "Thank"})

	fmt.Println(^uintptr(0) == math.MaxUint64)
}
