package main

import (
	"fmt"
	"strconv"
)

func main() {
	// string 转换为bool
	var str1 string = "true"
	var b bool
	b, _ = strconv.ParseBool(str1)
	fmt.Printf("b type is %T，b value is %v\n", b, b)
	// string 转换为 int
	var str2 string = "1023"
	var num2 int64
	var num3 int
	num2, _ = strconv.ParseInt(str2, 10, 64)
	num3 = int(num2)
	fmt.Printf("num3 type is %T，num3 value is %v\n", num3, num3)
	// string 转换浮点型
	var str3 string = "23.43"
	var c float64
	c, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("c type is %T，c value is %v\n", c, c)
}
