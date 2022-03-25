package main

import "fmt"

func main() {
	var LineNumber int
	fmt.Println("请输入你要答应的行数：\n")
	fmt.Scanf("%d", &LineNumber)
	fmt.Println("------------------\n")
	for i := 0; i <= LineNumber; i++ {
		for j := LineNumber; j >= 0; j-- {
			if i+j == LineNumber {
				fmt.Printf("%d  +  %d  =  %d\n", i, j, i+j)
			}
		}
	}
	fmt.Println("------------------")
}
