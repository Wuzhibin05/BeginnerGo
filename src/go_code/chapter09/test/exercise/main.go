package main

import "fmt"

func modifyUser(users map[string]map[string]string, name string) {
	_, ok := users[name] // if users[name] !=nil
	if ok {
		users[name]["pwd"] = "888888"
	} else {
		users[name] = make(map[string]string)
		users[name]["pwd"] = "9999"
		users[name]["nickname"] = "bbb"
	}
}

func main() {
	//定义用户map,存入基本信息
	var userMap = make(map[string]map[string]string)
	userMap["smith"] = make(map[string]string, 2)
	userMap["smith"]["age"] = "30"
	userMap["smith"]["pwd"] = "999"
	userMap["smith"]["nickname"] = "aaa"

	modifyUser(userMap, "tom")
	modifyUser(userMap, "smith")
	modifyUser(userMap, "alice")
	fmt.Println(userMap)
}
