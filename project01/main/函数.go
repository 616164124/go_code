package main

import "fmt"

func main() {
	i := max(12, 21)
	fmt.Println(i)
}

func max(a int, b int) int {
	//局部变量
	if a < b {
		return b
	} else {
		return a
	}
}
