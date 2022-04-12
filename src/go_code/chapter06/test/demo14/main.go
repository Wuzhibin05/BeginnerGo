package main

import (
	"fmt"
	"time"
)

func main() {
	for n := 0; n < 10; n++ {
		time.Sleep(time.Millisecond * 1000)
		fmt.Println("n=", n)
	}
}
