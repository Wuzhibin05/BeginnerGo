package main

import (
	"fmt"
)

type MethodUtils struct {
}

func (Met MethodUtils) matrix(m int, n int, str string) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(str)
		}
		fmt.Println()
	}
}

func main() {
	var m1 MethodUtils
	//m1.matrix()
	m1.matrix(2, 3, "#")
}
