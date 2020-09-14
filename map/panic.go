package main

import (
	"fmt"
)

func main() {
	var result = make(map[int]int)
	go func(){
		defer func(){
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()

		for{
			_ = result[1]
		}
	}()

	for{
		result[1]=1
	}
	fmt.Println("process over")
}