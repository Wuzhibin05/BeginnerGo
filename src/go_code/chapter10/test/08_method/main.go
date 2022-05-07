package main

import "fmt"

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c *Circle) area1() float64 {
	c.radius = 10.0
	return 3.14 * c.radius * c.radius
}

func main() {
	var circle1 = Circle{
		radius: 3.22,
	}
	res := circle1.area()
	fmt.Printf("圆的面积为：%.2f\n", res)

	res1 := circle1.area()
	fmt.Printf("圆的面积为：%.2f", res1)
}
