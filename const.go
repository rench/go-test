package main

import (
	"fmt"
	"unsafe"
)

//姓名
const (
	Unkonwn = 0
	Female  = 1
	Male    = 2
)

const (
	x = "xyz"
	y = len(x)
	z = unsafe.Sizeof(x)
)

func main() {
	const LENGTH = 12
	const WIDTH = 5

	fmt.Println(LENGTH * WIDTH)

	var a, b, c = 1, "is", true
	fmt.Println(a, b, c)

	fmt.Println("Male is ", Male)

	fmt.Println(a, b, c)
	var aa = &a

	fmt.Println("", aa)

}
