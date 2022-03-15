//1. 创建一个基于 for 的简单的循环。使其循环 10 次，并且使用 fmt 包打印出计数 器的值
package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("i的值为：%d\n", i)
	}
}
