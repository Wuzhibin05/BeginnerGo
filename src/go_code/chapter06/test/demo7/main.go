package main

import "fmt"

func main() {

	a := func(x int, y int) int {
		return x + y
	}
	fmt.Println(a(10, 20))
	fmt.Println(a(11, 20))
	fmt.Println(a(12, 20))
}
