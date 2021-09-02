package main

import "fmt"

func main() {
	var aMap map[string]string
	aMap = make(map[string]string)
	aMap["a"] = "111111"
	aMap["b"] = "2222"
	aMap["c"] = "3333"
	aMap["d"] = "444"
	for _, a := range aMap {
		fmt.Println(a)
	}
	//aa 表示的为key为a的value值，ok为时候存在
	aa, ok := aMap["a"]
	if ok {
		fmt.Println("存在", aa)
	} else {
		fmt.Println("不在", aa)
	}
}
