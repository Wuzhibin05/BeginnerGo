package main

import (
	"fmt"
	"sort"
)

func main() {

	//map的排序
	map1 := make(map[int]int)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90

	fmt.Println(map1)
	var keys []int
	fmt.Printf("type=%T,len=%v,cap=%v\n", keys, len(keys), cap(keys))
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	fmt.Println(keys)
	sort.Ints(keys)
	fmt.Println(keys)
	for _, value := range keys {
		fmt.Printf("k=%v，v=%v\n", value, map1[value])
	}
}
