package main

import (
	"fmt"
)

type MethodUtils struct {
}

func (m MethodUtils) matrix() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (Met MethodUtils) matrix1(m, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	var m1 MethodUtils
	//m1.matrix()
	m1.matrix1(2, 3)
}
