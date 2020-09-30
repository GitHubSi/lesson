package main

import "fmt"

func main() {
	var ch chan int = make(chan int, 3)
	ch <- 1
	ch <- 2
	close(ch)

	// 第一次读
	f, ok := <-ch
	fmt.Printf("ok is %t, value is %d\n", ok, f)

	// 第二次读
	s, ok := <-ch
	fmt.Printf("ok is %t, value is %d\n", ok, s)

	// 第三次读
	t, ok := <-ch
	fmt.Printf("ok is %t, value is %d\n", ok, t)
}
