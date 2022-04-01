package utils

import "fmt"

var (
	Age  byte
	Name string
)

func init() {
	fmt.Println("In Utils init() start ...")
	Age = 33
	Name = "Wuzhibin"
}
