package main

import "fmt"

func AddUpper() func(int) int {
	var n int = 10
	return func(i int) int {
		n = n + i
		return n
	}
}

func main() {
	// f = AddUpper() 相当于调用函数AddUpper，而他的返回值为 fun(int)int的匿名函数
	// f = fun(int)int
	// f(1) = func (1)  这个函数没有名字，实际应该没有func,就是调用匿名函数传入1的时候的返回值
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	fmt.Println(f(4))
}
