package main

import "fmt"

func main() {
	var sum int
	var total int
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			sum += 1
			total += i
			fmt.Printf("i=%d\n", i)
		}
	}
	fmt.Printf("sum=%d,total=%d", sum, total)
}
