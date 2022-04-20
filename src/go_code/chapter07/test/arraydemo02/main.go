package main

import "fmt"

func main() {
	//从终端循环输入 5 个成绩， 保存到 float64 数组,并输出.
	var score [5]float64
	num := len(score)
	for i := 0; i < num; i++ {
		fmt.Printf("请输入第%v个学生的成绩:\n", i+1)
		fmt.Scan(&score[i])
	}
	for index, value := range score {
		fmt.Printf("第%d个学生的成绩为：%.2f\n", index, value)
	}
}
