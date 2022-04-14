package main

import "fmt"

func fbn(n int) []uint64 {
	var fbnSlice = make([]uint64, n)
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}

func main() {
	res := fbn(10)
	fmt.Println("res=", res)
}
