package main

import "fmt"

func main() {
	var maps map[string]string
	maps = make(map[string]string)
	maps["A"] = "香蕉"
	maps["B"] = "苹果"
	fmt.Printf("maps存放地址:%p\r\n", &maps)
	fmt.Printf("maps type:%T\r\n", maps)
	for k, v := range maps {
		fmt.Printf("%s @ %p - %s @ %p\r\n", k, &k, v, &v)
	}
	fmt.Println("=====================")
	addFruit("C", "橘子", &maps)
	for k, v := range maps {
		fmt.Printf("%s @ %p - %s @ %p\r\n", k, &k, v, &v)
	}
	fmt.Println("=====================")
	addFruit2("D", "火龙果", maps)
	for k, v := range maps {
		fmt.Printf("%s @ %p - %s @ %p\r\n", k, &k, v, &v)
	}

}

func addFruit(name string, value string, maps *map[string]string) {
	var tmp = *maps
	tmp[name] = value
	fmt.Printf("指针存放地址:%p\r\n", &maps)
	fmt.Printf("*maps存放地址:%p\r\n", *maps)
	fmt.Printf("maps存放地址:%p\r\n", maps)
}

func addFruit2(name string, value string, maps map[string]string) {
	maps[name] = value
	fmt.Printf("maps存放地址:%p\r\n", &maps)
	fmt.Printf("maps存放地址2:%p\r\n", maps)
	fmt.Printf("maps存放地址2s:%s\r\n", maps)
	fmt.Printf("maps存放地址2s:%p\r\n", *&maps)
}
