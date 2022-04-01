package main

import "fmt"

var (
	fun1 = func(x int, y int) int {
		return x + y
	}
)

func main() {
	fmt.Println(fun1(10, 20))
}
