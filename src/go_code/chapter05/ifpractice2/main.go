package main

import "fmt"

func main() {
	var month byte
	var age byte
	var price int = 60

	fmt.Println("请输入游玩的月份:")
	fmt.Scanln(&month)
	fmt.Println("请输入你的年龄:")
	fmt.Scanln(&age)
	if month >= 4 && month <= 10 {
		if age < 18 {
			fmt.Println("你的票价为：", price/2)
		} else if age > 60 {
			fmt.Println("你的票价为：", price/3)
		} else {
			fmt.Println("你的票价为：", price)
		}
	} else {
		if age >= 18 && age <= 60 {
			fmt.Println("你的票价为：40")
		} else {
			fmt.Println("你的票价为：20")
		}
	}
}
