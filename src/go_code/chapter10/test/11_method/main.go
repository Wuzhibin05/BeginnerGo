package main

import (
	"fmt"
)

type rectangle struct {
	len   float64
	width float64
}

func (r rectangle) area() float64 {
	return r.len * r.width
}

func main() {
	rectangle1 := rectangle{
		len:   3.12,
		width: 12.12,
	}
	fmt.Printf("矩形的面积：%.2f", rectangle1.area())
}
