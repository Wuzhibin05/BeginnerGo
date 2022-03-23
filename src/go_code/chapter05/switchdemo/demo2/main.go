package main

import "fmt"

func main() {
	var x interface{}
	var y = 10.22
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型~：%T", i)
	case int:
		fmt.Printf("x是int类型")
	case float64:
		fmt.Printf("x是float64类型")
	case func(int) float64:
		fmt.Printf("x是func(int)类型")
	case string, bool:
		fmt.Printf("x是string或者bool类型")
	default:
		fmt.Printf("x是未知类型")
	}
}
