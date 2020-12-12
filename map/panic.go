package main

import (
	"fmt"
	"time"
)

type Value struct {
	A string
	B int
	C time.Time
	D []byte
	E float32
	F *string
	G People
}
type People struct {
	Sex     int64
	Age     int64
	Name    string
	Icon    string
	Address string
	Height  int64
}

func main() {
	a := []string{"a"}
	fmt.Println(a)

	for{

	}
}
