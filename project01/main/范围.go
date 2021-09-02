package main

import "fmt"

func main() {
	nums := []int{1, 2, 5, 11, 23, 6, 100}
	sum := 0
	nums1 := []int{1, 2, 5, 11, 23, 6, 100}
	sum1 := 0
	s := "wwww"
	for _, num := range nums {
		sum += num
	}

	fmt.Println(sum, s)
	/*******************************************/
	//i 表示下标
	for i, num1 := range nums1 {
		sum1 += num1
		fmt.Println(sum1, i)
	}
	fmt.Println(sum1)
//

}
