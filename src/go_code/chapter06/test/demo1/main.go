package main

import "fmt"

func test(n1 int) int {
	n1 = n1 + 1
	fmt.Println("test(), n1=", n1)
	return n1
}

func getSum(n1 int, n2 int) int {
	var sum int
	sum = n1 + n2
	fmt.Println("getSum sum=", sum)
	return sum
}

func main() {
	n1 := 10
	n2 := 20
	test(n1)
	fmt.Println("main(),n1=", n1)
	getSum(n1, n2)
}
