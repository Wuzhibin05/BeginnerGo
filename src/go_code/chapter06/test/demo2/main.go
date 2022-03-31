package main

import "fmt"

func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {
	n1 := 10
	n2 := 20

	res1, _ := getSumAndSub(n1, n2)
	fmt.Println("sum", res1)
	_, res2 := getSumAndSub(n1, n2)
	fmt.Println("sub", res2)
}
