package main

import (
	"fmt"
	"sync"
	// "time"
)

//采用sync.WaitGroup的方式来控制
/*
在主线程中，开启一个goroutine，该协程每隔50毫秒输出，“你好，kk”
在主线程中也每隔一秒输出，“你好，tt”，输出10次后退出程序
要求goroutine与主线程同事执行
*/
func test() {
	for i := 0; i < 200000; i++ {
		fmt.Println("你好，kk--", i)
	}
	wg.Done()
}
//全局变量
var wg sync.WaitGroup
func main() {
	wg.Add(1)
	go test()
	for i := 0; i < 5; i++ {
		fmt.Println("你好，tt---", i)
	}
	wg.Wait()
}
