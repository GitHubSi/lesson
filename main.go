package main

import (
	"fmt"
	"strconv"
	"sync"
)

type Binary uint64

func (i Binary) String() string {
	return strconv.FormatUint(i.Get(), 10)
}

func (i Binary) Get() uint64 {
	return uint64(i)
}

type L struct {
	m sync.RWMutex
}

func main() {
	fmt.Println(1 << 0)
	l := L{}
	l.m.Lock()
	//l.m.Unlock()
	fmt.Println("thx")
}
