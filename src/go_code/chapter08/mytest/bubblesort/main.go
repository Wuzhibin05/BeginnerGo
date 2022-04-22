package main

import "fmt"

func main() {
	var array1 = [5]int{24, 69, 80, 57, 13}
	fmt.Println(array1)
	for i := 0; i < (len(array1) - 1); i++ {
		for j := 0; j < (len(array1) - 1 - i); j++ {
			if array1[j] > array1[j+1] {
				t := array1[j+1]
				array1[j+1] = array1[j]
				array1[j] = t
			}
		}
	}
	fmt.Println(array1)
}
