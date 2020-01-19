package main

import "fmt"

func main() {
	age := 12
	name := "张三"

	if age > 18 {
		fmt.Println(name + "成年!")
	} else {
		fmt.Printf(name + "未成年!")
	}

	switch age {
	case 12:
		fmt.Println(name + "刚满12岁!")
	default:
		fmt.Println(name+"年龄", age)
	}

}
