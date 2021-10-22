package main

import (
	"fmt"
)
//main执行完，test协程就结束了（不管有没有完全执行完成）
/*
在主线程中，开启一个goroutine，该协程每隔50毫秒输出，“你好，kk”
在主线程中也每隔一秒输出，“你好，tt”，输出10次后退出程序
要求goroutine与主线程同事执行
*/
func test(){
	for i := 0 ;i<200 ; i++{
		fmt.Println("你好，kk--" ,i )
	}

}

func main(){
	go test()
	for i := 0; i < 2; i++ {
		fmt.Println("你好，tt---",i)
	}
	

}