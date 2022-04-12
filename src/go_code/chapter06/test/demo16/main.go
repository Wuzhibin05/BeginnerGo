// 编写一段代码，测试test3的执行时间
package main

import (
	"fmt"
	"time"
)

func test3() {
	time.Sleep(time.Second * 5)
	fmt.Println("test03()执行完毕")
}

func main() {
	startTime := time.Now().UnixNano()
	test3()
	endTime := time.Now().UnixNano()
	timeDuration := endTime - startTime
	fmt.Printf("test03执行时间为：%v纳秒", timeDuration)
}
