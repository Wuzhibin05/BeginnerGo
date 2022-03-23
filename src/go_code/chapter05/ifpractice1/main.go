package main

import "fmt"

func main() {
	var grade float64

	fmt.Println("请输入你的百米成绩（秒）:")
	fmt.Scanln(&grade)
	if grade < 8.0 {
		var gender string
		fmt.Println("请输入你的性别（男，女）:")
		fmt.Scanln(&gender)
		if gender == "男" {
			fmt.Println("你成功进入男子组决赛！")
		} else {
			fmt.Println("你成功进入女子组决赛！")
		}
	} else {
		fmt.Println("你的成绩太差，已经被淘汰！")
	}
}
