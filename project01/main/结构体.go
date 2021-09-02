package main

import "fmt"

type Book struct {
	name   string
	author string
	id     int
}

func main() {
	var book1 Book
	
	//创建Book 方式一 
	fmt.Println(Book{"世界","xioa",1990})

	//创建Book (key => value) 方式二
	fmt.Println(Book{name:"世界",author: "xioa",id: 123})
	fmt.Println(Book{name:"世界",id: 123})
	book1.author="jfks"
	book1.id=123
	book1.name="kjfks"
	fmt.Println(book1)
	var book2 *Book
	book2=&book1
	book2.author="bbbk2"
	fmt.Printf("book2.author: %v\n", book2)
	


}