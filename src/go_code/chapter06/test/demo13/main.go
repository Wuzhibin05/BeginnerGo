package main

import (
	"errors"
	"fmt"
)

//函数去读取以配置文件init.conf的信息
//如果文件名传入不正确，我们就返回一个自定义的错误
func readconf(name string) (err error) {
	if name == "init.conf" {
		return nil
	} else {
		return errors.New("读取文件错误")
	}
}

func test02() {
	err := readconf("init1.conf")
	if err != nil {
		panic(err)
	}
	fmt.Println("test02继续执行")
}

func main() {
	// 测试自定义错误
	test02()
	fmt.Println("main()下面的代码...")
}
