package main

import "fmt"

type monstor struct {
	Name  string
	Age   int
	Skill string
}

func main() {
	var monstor1 monstor
	monstor1.Name = "狮子精"
	monstor1.Age = 20

	monstor2 := monstor1
	monstor2.Name = "玉兔精"
	monstor2.Age = 30

	fmt.Printf("monstor1的名字：%v,年龄:%v\n", monstor1.Name, monstor1.Age)
	fmt.Printf("monstor2的名字：%v,年龄:%v", monstor2.Name, monstor2.Age)
}
