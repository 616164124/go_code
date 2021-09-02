package main

import "fmt"

func main() {
	const a string = "iiii" //const 表示常量，后续不能重新赋值，整个程序中可以不使用到
	var b string = "iiii"   //var 后续可以重新赋值，但是整个程序中必须使用到
	b = "aaaa"
	fmt.Println(a, b)
	const e, f, g = 1, false, "sss" // 多重赋值
	const aa, bb, cc = 2, 3, 4      //多重赋值
	var dd int32
	dd = aa * bb // 实际为 2 * 3 =6
	fmt.Println("有多大%D", dd)
	//如果数组个数已经被限制，那么定义时不能超出这个限制
	var slice1 = []float64{12.1, 3.4, 8.3}

	var slice2 = [4]string{"1", "23", "3423", "3243"}

	fmt.Println(slice1[2])
	fmt.Println(slice2)

	for i := 0; i < 3; i++ {
		fmt.Println(&slice1[i])
	}
	/*******************************/
	var slice1 []type= make([]type, len)
	fmt.Println(slice1)
}
