package main

import (
	"fmt"
	"time"
)

//加上time来实现控制(不能完全知道具体时间长短)
/*
在主线程中，开启一个goroutine，该协程每隔50毫秒输出，“你好，kk”
在主线程中也每隔一秒输出，“你好，tt”，输出10次后退出程序
要求goroutine与主线程同事执行
*/
func test(){
	for i := 0 ;i<200 ; i++{
		fmt.Println("你好，kk--" ,i )
		time.Sleep(time.Millisecond *1)
	}

}

func main(){
	go test()
	for i := 0; i < 5; i++ {
		fmt.Println("你好，tt---",i)
		time.Sleep(time.Millisecond *1)
	}

}