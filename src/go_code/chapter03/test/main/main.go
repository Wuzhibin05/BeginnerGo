package main

import "fmt"

func main() {
	var a int = 300
	var ptr *int = &a // 错误，变量类型不匹配
	fmt.Println(ptr)
}
