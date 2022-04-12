package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("now=%v,tpye=%T\n", now, now)
	fmt.Println(now.Year())
	fmt.Println(now.Day())
	fmt.Println(now.Month())
	fmt.Println(now.Hour())

	// 方式一：
	fmt.Printf("%d/%d/%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	// Sprintf生成格式化字符串并返回该字符串
	dateStr := fmt.Sprintf("%d/%d/%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Print(dateStr)
	// 方式二：使用time.format
	dateStr2 := now.Format("2006-01-02 15:04:05")
	fmt.Println(dateStr2)
}
