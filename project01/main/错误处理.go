package main

import (
	"errors"
	"fmt"
)

func main() {

	i := error1(21, 12)
	fmt.Println(i)

}

func error1(a int, b int) {
	if a<b {
		return 11,errors.New("cuowu!!!!")
	}else {
		// return 12,errors.New("错误！！！！！")
	}

	// return 1
}