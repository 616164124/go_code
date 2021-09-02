package main

// 参考网址 https://bbs.huaweicloud.com/blogs/222284
import (
	"fmt"
	"strconv"
)

func main() {
	var sum int16 = 12
	var sum1 int16 = 13
	fmt.Println(float64(sum) / float64(sum1))
	v := 123
	vS := strconv.Itoa(v)
	fmt.Println(vS)
	var vi64 int64 = 7663
	//10表示为十进制
	vInt64 := strconv.FormatInt(vi64, 10)
	fmt.Println((vInt64))
}
