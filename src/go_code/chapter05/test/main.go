package main

import "fmt"

func main() {
	var height int32
	var money float32
	var handsome bool
	fmt.Println("请输入你的身高(厘米):")
	fmt.Scanln(&height)
	fmt.Println("请输入你的资产（千万）:")
	fmt.Scanln(&money)
	fmt.Println("请输入你是否长得帅（true/fasle）:")
	fmt.Scanln(&handsome)
	if height > 180 && money > 1 && handsome == true {
		fmt.Println("我一定要嫁给他")
	} else if height > 180 || money > 1 || handsome == true {
		fmt.Println("嫁吧，比上不足，比下有余")
	} else {
		fmt.Println("坚决不嫁")
	}
}
