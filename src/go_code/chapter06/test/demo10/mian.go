package main

import (
	"fmt"
)

var (
	name = "Tom"
)

func test1() {
	fmt.Println(name)
}
func test2() {
	name := "jack"
	fmt.Println(name)
}
func main() {
	fmt.Println(name) //Tom
	test1()           // Tom
	test2()           // jack
	test1()           // tom
}
