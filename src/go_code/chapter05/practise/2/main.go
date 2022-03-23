package main

import "fmt"

func main() {
	var score float64
	fmt.Println("请输入学生成绩：")
	fmt.Scanf("%f", &score)
	if score <= 100.0 && score >= 0.0 {
		switch score >= 60.0 {
		case true:
			fmt.Println("成绩合格")
		default:
			fmt.Println("成绩不合格")
		}
	} else {
		fmt.Println("输入的成绩不合法")
	}
}
