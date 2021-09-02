package main

import "fmt"

func main() {
	var a int = 22
	var ip *int
	ip = &a
	fmt.Println(&a)
	fmt.Println(ip)
	//使用指针访问值
	fmt.Println(*ip)
	var bb byte = 'a'
	var zhizhen *byte
	zhizhen = &bb
	fmt.Println(*zhizhen) //输出的是a的asci值
	var c *int
	if c!=nil {
		fmt.Println("c不为nil")
	}else{
		fmt.Println("c为nil")
	}
	fmt.Println(c) //nil 为空指针

}
