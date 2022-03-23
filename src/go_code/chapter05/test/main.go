package main

import "fmt"

func main() {
	var week string
	fmt.Println("请输入一个字符（a,b,c,d,e,f,g）:")
	fmt.Scanln(&week)
	switch week {
	case "a":
		fmt.Println("周一猴子穿新衣。")
	case "b":
		fmt.Println("周二猴子穿新衣。")
	case "c":
		fmt.Println("周三猴子穿新衣。")
	case "d":
		fmt.Println("周四猴子穿新衣。")
	case "e":
		fmt.Println("周五猴子穿新衣。")
	case "f":
		fmt.Println("周六猴子穿新衣。")
	case "g":
		fmt.Println("周日猴子穿新衣。")
	default:
		fmt.Println("输入的字符不在要求的范围内！")
	}
}
