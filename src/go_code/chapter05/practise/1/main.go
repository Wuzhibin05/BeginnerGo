package main

import "fmt"

func main() {
	var n1 byte
	fmt.Println("请输入你要转换的小写字符（例如a,b，c...）")
	fmt.Scanf("%c", &n1)
	switch n1 {
	case 'a', 'b', 'c', 'd', 'e':
		fmt.Printf("转换后的字符为%c", n1-32)
	default:
		fmt.Println("other")
	}
}
