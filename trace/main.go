package main

import (
	"github.com/think-next/lesson/trace/trace"
	"sync"
	"time"
)

func A1() {
	defer trace.Trace()()
	B1()
}

func B1() {
	defer trace.Trace()()
	C1()
}

func C1() {
	defer trace.Trace()()
	D()
}

func D() {
	defer trace.Trace()()
}

func A2() {
	defer trace.Trace()()
	B2()
}
func B2() {
	defer trace.Trace()()
	C2()
}
func C2() {
	defer trace.Trace()()
	D()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		A2()
		wg.Done()

		go func() {
			A2()
			wg.Done()
		}()
	}()

	time.Sleep(time.Millisecond * 50)
	A1()
	wg.Wait()
}
