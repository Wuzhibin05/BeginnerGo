package main

import "fmt"

func pyramid(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= num-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println() // 默认打印换行
	}
}

func main() {
	var layer int
	fmt.Println("请输入要打印的金字塔的层数（数字）：")
	fmt.Scanln(&layer)
	pyramid(layer)
}
