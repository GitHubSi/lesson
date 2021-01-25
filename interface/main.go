package main

import (
	"fmt"
	"github.com/think-next/lesson/interface/inner"
)

type Point struct {
	X int
	Y int
}

type DescPoint struct {
	Point // 匿名结构体
	Desc  string

	// 匿名接口
	// inner.Inter
}

func (desc *DescPoint) Type() int {
	return 2
}

func (desc *DescPoint) name() string {
	return ""
}

func main() {
	desc := &DescPoint{
		Point: Point{},
		Desc:  "empty",
	}

	desc.Point.X = 1
	desc.Y = 2

	var i inner.Inter = desc
	fmt.Println(i.Type())
}
