// 3.再次改写这个循环，使其遍历一个 array，并将这个 array 打印到屏幕上
package main

import (
	"fmt"
)

func main() {
	var a = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(a); i++ {
		fmt.Printf("a数组中第%d个元素的值为：%d\n", i, a[i])
	}
}
