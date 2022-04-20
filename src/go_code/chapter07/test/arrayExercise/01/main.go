package main

import (
	"fmt"
)

func main() {
	var alphabet [26]byte
	alphabet[0] = 'A'
	for i := 1; i < len(alphabet); i++ {
		alphabet[i] = alphabet[i-1] + 1
	}
	for index, value := range alphabet {
		fmt.Printf("第%d个字母是%c\n", index, value)
	}
}
