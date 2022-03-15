//2. 用 goto 改写 1 的循环。关键字 for不可使用。

package main

import (
	"fmt"
)

func main() {
	var i int = 1
LOOP:
	fmt.Printf("i的值为：%d\n", i)
	if i < 10 {
		i++
		goto LOOP
	}
}
