package main

import "fmt"

func main() {
	//1 创建channel
	ch := make(chan int, 3)
	//2 给管道里面存储内容(不能超出创建时的容量length)
	ch <- 12
	ch <- 10
	ch <- 9
	// ch <- 100
	//3 获取管道内的内容
	/*
		每次 <-ch 就输出一个值，
	*/
	a := <-ch
	<-ch
	b := <-ch
	fmt.Println(a)
	fmt.Println(b)
	ch <- 888 //又给channel添加一个值
	//4 打印管道的长度和容量
	fmt.Println("值：%v 容量，%v, 长度 %v \n", ch, cap(ch), len(ch))
	// 5 管道的类型（引用类型）

	ch1 := make(chan int, 4)
	ch1 <- 11
	ch1 <- 21
	ch1 <- 33
	ch2 := ch1
	ch2 <- 55
	fmt.Println(len(ch1))
	<-ch1
	<-ch1
	<-ch1
	d := <-ch1
	fmt.Println(d)  //55 

//7 管道阻塞
	ch3:=make(chan int ,1)
	ch3<-12
	ch3<-1  // all goroutines are asleep - deadlock!   阻塞
	

}
