package main

import "fmt"

func main() {
Lable:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				break Lable
			}
			fmt.Println("j=", j)
		}
	}
}
