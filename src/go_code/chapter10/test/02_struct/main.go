package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Score  [6]float64
	Ptr    *int
	Slice1 []int
	Map1   map[string]string
}

func main() {
	num := 60
	var p1 Person
	fmt.Println("Name=", p1.Name)
	fmt.Println("Age=", p1.Age)
	fmt.Println("Score=", p1.Score)
	fmt.Println("Ptr=", p1.Ptr)
	fmt.Println("Slice1=", p1.Slice1)
	fmt.Println("Map1=", p1.Map1)
	// 默认都是nil，指针，切片，还有map
	if p1.Ptr == nil {
		fmt.Println("ok1")
	}
	if p1.Slice1 == nil {
		fmt.Println("ok2")
	}
	if p1.Map1 == nil {
		fmt.Println("ok3")
	}
	p1.Slice1 = make([]int, 6)
	p1.Map1 = map[string]string{
		"name":  "鲁智深",
		"hobby": "喝酒",
	}
	p1.Ptr = &num
	fmt.Println("Ptr=", *(p1.Ptr))
	fmt.Println("Slice1=", p1.Slice1)
	fmt.Println("Map1=", p1.Map1)

}
