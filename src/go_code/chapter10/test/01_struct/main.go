package main

import "fmt"

type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
}

func main() {
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	cat1.Hobby = "吃饭和睡觉"

	var cat2 Cat
	cat2.Name = "小黑"
	cat2.Age = 100
	cat2.Color = "黑色"
	cat2.Hobby = "跳舞"

	fmt.Println("猫咪的信息如下：\n", cat1)
	fmt.Println("name=", cat1.Name)
	fmt.Println("Age=", cat1.Age)
	fmt.Println("Color=", cat1.Color)
	fmt.Println("Hobby=", cat1.Hobby)
	fmt.Println("猫咪的信息如下：\n", cat2)
	fmt.Printf("%p", &cat1)

}
