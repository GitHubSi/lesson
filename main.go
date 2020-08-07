package main

import (
	"strconv"
)

type Binary uint64

func (i Binary) String() string {
	return strconv.FormatUint(i.Get(), 10)
}

func (i Binary) Get() uint64 {
	return uint64(i)
}

func main() {

}
