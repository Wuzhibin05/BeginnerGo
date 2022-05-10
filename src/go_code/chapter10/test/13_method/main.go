package main

import "fmt"

type MethodUtils struct {
}

func (Met MethodUtils) multiTable(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v ", j, i, i*j)
		}
		fmt.Print("\n")
	}
}

func main() {
	var table MethodUtils
	//m1.matrix()
	table.multiTable(9)
}
