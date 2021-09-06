package main

import "fmt"

func main() {
	i := max(12, 21)
	fmt.Println(i)
}
//	func 函数名(参数1 参数类型,参数2 参数类型) 返回值类型{}
func max(a int32, b int32) int {
	//局部变量
	if a < b {
		return b
	} else {
		return a
	}
}
