package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) jisuan() {
	res := 0
	for i := 1; i <= 1000; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算结果为：", res)
}

func (p Person) jisuan1(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func (p Person) GetSum(m, n int) int {
	return m + n
}

func (p Person) speak() {
	fmt.Printf("%v是一个好人!\n", p.Name)
}

// 绑定方法 test
func (p Person) test() {
	fmt.Println("test().....,name=", p.Name)
}

func main() {
	var p Person
	p.Name = "Tom"
	p.test() // 调用方法
	p.speak()
	p.jisuan()
	res := p.GetSum(100, 10)
	fmt.Println("res=", res)

	res1 := p.jisuan1(100)
	fmt.Println("res=", res1)
}
