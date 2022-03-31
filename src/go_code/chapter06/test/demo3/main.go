// 编写一个函数，可以求一个数和多个数的和
package main

import "fmt"

func getSum(n int, args ...int) int {
	var sum int
	sum = n
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func main() {
	res := getSum(10, 1, 2, 4, 5, 5, 6)
	fmt.Println("res=", res)
}
