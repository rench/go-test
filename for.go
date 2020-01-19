package main

import (
	"fmt"
	"strconv"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var names = []string{"hello", "world"}
	for i, name := range names {
		fmt.Printf("第 %d 位是 %s\n", i, name)
	}

	var arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(arr[4:])
	var num = 1.03
	fmt.Println(strconv.FormatFloat(num, 'E', -1, 64))
	fmt.Println(strconv.FormatFloat(num, 'f', 6, 64))

	var numstr = "1.03"
	var tt, err = strconv.ParseFloat(numstr, 64)
	fmt.Println(tt, err)

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
