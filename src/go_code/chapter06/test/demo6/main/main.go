package main

import (
	"fmt"
	"go_code/chapter06/test/demo6/utils"
)

var age = test()

func test() int {
	fmt.Println("test() start...")
	return 90
}

func init() {
	fmt.Println("init() start....")

}

func main() {
	fmt.Println("main() start....")
	fmt.Printf("name=%v,age=%v", utils.Age, utils.Name)
	fmt.Println(age)
}
