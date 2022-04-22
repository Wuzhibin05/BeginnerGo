package main

import "fmt"

func main() {
	// 方式一：直接引用数组
	var array1 = [7]int{1, 2, 43, 22, 12, 4234, 2131}
	slice1 := array1[1:5]
	fmt.Println("slice1=", slice1)

	// 方式二：通过make来定义
	var slice2 []int = make([]int, 10)
	fmt.Println(slice2)

	// 方式三：直接定义，类似定义数组，不指定长度即可
	var slice3 = []int{1, 2, 43, 2141}
	fmt.Printf("slice3的类型是%T,len=%v,cap=%v", slice3, len(slice3), cap(slice3))
}
