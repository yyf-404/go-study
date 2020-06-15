package main

import (
	"fmt"
	"strconv"
)

type Book struct {
	title string
	author string
	subject string
	book_id int
}
/* 通过自定义工厂方法来模拟构造函数 */
func createBook(title string,author string,subject string,book_id int) *Book{
	return &Book{title,author,subject,book_id}
}
func main() {
	// 创建一个新的结构体
	fmt.Println(Book{"Go", "yyf", "Go语言教程", 6435})
	// 也可以使用 key => value 格式
	fmt.Println(Book{title: "Go", book_id: 6435, author: "yyf", subject: "Go语言教程"})
	// 忽略的字段为 0 或 空
	fmt.Println(Book{title: "Go", author: "yyf"})
	var Book2 Book //申明
	Book2=Book{title: "Java", book_id: 9326, author: "yyf", subject: "Java语言教程"} //赋值
	fmt.Println(Book2)
	var Book1 Book        /* 声明 Book1 为 Book 类型 */
    Book1.title="JAVA"
    Book1.book_id=9326
    fmt.Println("Book 1 title : "+Book1.title)
	fmt.Println("Book 1 title : "+strconv.Itoa(Book1.book_id)) //int转字符
	printBook(Book{title: "Go", book_id: 6435, author: "yyf", subject: "Go语言教程"})
    var ptr *Book
    ptr=&Book2
    fmt.Println("Book 2 title : "+ptr.title)
    fmt.Println(createBook("cpp", "yyf", "cpp语言教程", 6135))
}
func printBook(book Book) {
	fmt.Println("Book 1 title : "+book.title)
	fmt.Println("Book 1 title : "+strconv.Itoa(book.book_id)) //int转字符
}
