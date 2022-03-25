package main

import "fmt"

func main() {
	var avg0 float64
	var total float64
	for i := 0; i < 3; i++ {
		sum := 0.0
		avg1 := 0.0
		for j := 1; j < 6; j++ {
			var score float64 = 0.0
			fmt.Printf("请输入第%d班的第%d孩子的成绩:", i+1, j+1)
			fmt.Scanf("%d", &score)
			sum += score
		}
		avg1 = sum / 5
		total += sum
		fmt.Printf("第%d个班的平均分为：%f", i, avg1)
	}
	avg0 = total / 15
	fmt.Printf("三个班的平均分为：%f", avg0)
}
