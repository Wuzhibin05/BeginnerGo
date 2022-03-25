package main

import "fmt"

func main() {
	var money float64 = 100000
	var num int
	for {
		if money < 1000 {
			break
		}
		num++
		if money > 50000 {
			money -= money * 0.05
			continue
		}
		money -= 1000

	}
	fmt.Printf("通过关卡的数量是：%d,剩余金额为：%f", num, money)
}
