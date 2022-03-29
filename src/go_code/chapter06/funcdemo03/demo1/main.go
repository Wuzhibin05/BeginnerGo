package main

import (
	"fmt"
)

func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("n=", n)
}
func main() {

	//看一段代码
	//test(4) // ?通过分析来看下递归调用的特点
	test(4)
}
