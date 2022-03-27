package main

import "fmt"

func main() {
	var user string
	var passWord string
	var tryTimes = 3
	for {
		fmt.Printf("您当前还剩%v次登录机会\n", tryTimes)
		fmt.Println("请输入登录用户名：")
		fmt.Scanln(&user)
		fmt.Println("请输入登录密码：")
		fmt.Scanln(&passWord)
		if user == "张无忌" && passWord == "888" {
			fmt.Printf("登录成功")
			break
		} else {
			tryTimes--
		}
		if tryTimes == 0 {
			fmt.Println("对不起，登录次数过多，已经退出")
			break
		}
	}
}
