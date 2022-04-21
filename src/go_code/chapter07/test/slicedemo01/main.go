package main

import "fmt"

func main() {
	var intarray = [6]int{10, 22, 34, 55, 209, 23}
	slice1 := intarray[1:4]
	slice1[0] = 99
	fmt.Println("intarray=", intarray)
	fmt.Println("slice1=", slice1)
	fmt.Println("slice1的长度", len(slice1))
	fmt.Println("slice1的容量", cap(slice1))
}
