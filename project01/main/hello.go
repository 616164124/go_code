package main

import (
	"fmt"
	"time"
)

func main() {
	var u int16
	u = 12324
	if u < 12 {
		fmt.Println("u<12")
	} else {
		fmt.Println("u>12")
	}
	var a float64 = 343.2
	a = 123
	fmt.Println(a)
	var Day = time.Now().Year()

	fmt.Println(Day)
	fmt.Println("11111")
	fmt.Print("123231")
	fmt.Println(u)

	//把hello中的第一个h字母换成汉字‘喝’
	var src = "hello"
	var string1 = "jflsjfjfslfj"+
	"flksnfdjsfkjs\""+
	"flksjfljflsfj"+
	"fjlksjfl"
	i := len(src)
	fmt.Println(string1)
	
	fmt.Println(i)
	var c = []rune(src)
	c[0] = '喝'
	s2 := string(c)
	
	fmt.Println(s2)
}
