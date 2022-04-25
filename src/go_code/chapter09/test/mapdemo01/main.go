package main

import "fmt"

func main() {
	var heros = make(map[string]string, 10)
	var hero1 map[string]string
	hero1 = make(map[string]string)
	heros["no1"] = "宋江"
	heros["no2"] = "卢俊义"
	heros["no3"] = "吴用"
	fmt.Println("heros=", heros)

	fmt.Println()
	hero1["no1"] = "宋江"
	hero1["no2"] = "卢俊义"
	hero1["no3"] = "吴用"
	fmt.Println("hero1=", hero1)
}
