//题 3:请编写一个函数 swap(n1 *int, n2 *int) 可以交换 n1 和 n2 的值
package main

import "fmt"

func swap(n1 *int, n2 *int) {
	*n1 = *n2 + *n1
	*n2 = *n1 - *n2
	*n1 = *n1 - *n2
}

func main() {
	n1 := 10
	n2 := 20
	swap(&n1, &n2)
	fmt.Printf("n1=%v,n2=%v", n1, n2)
}
