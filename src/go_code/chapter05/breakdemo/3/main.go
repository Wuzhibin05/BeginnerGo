package main

import "fmt"

func main() {
	var sum int = 0
	for i := 0; i < 100; i++ {
		sum += i
		if sum > 20 {
			fmt.Printf("i=%d，sum=%d", i, sum)
			break
		}
	}
}
