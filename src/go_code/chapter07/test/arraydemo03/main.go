package main

import "fmt"

func main() {
	//第一种方式：定义并赋值
	var array1 [3]int = [3]int{1, 2, 3}
	fmt.Println(array1)

	// 第二种方式：第一种方式的简化写法
	var array2 = [3]int{1, 2, 3}
	fmt.Println(array2)

	// 第三种：不指定长度，自动计算长度，用...代替（固定写法），第二种方式的简化写法
	var array3 = [...]int{1, 2, 3}
	fmt.Println(array3)

	// 第四种：不指定长度，通过下标来指定值
	var array4 = [...]int{1: 100, 3: 200, 0: 100}
	fmt.Println(array4)

	// 类型推导
	array5 := [...]string{1: "tom", 3: "jack", 0: "bruce"}
	fmt.Println(array5)
}
