package main

import "fmt"

func main() {
	/* 值传递 */
	val:=1
	fmt.Println("变量val地址为",&val)
	valueTransmit(val)
	fmt.Println("变量val值为",val)
	/* 指针传递*/
	var ptr *int
	ptr=&val
	fmt.Println("--------------------")
	pointerTransmit(ptr)
	fmt.Println("main函数中指针存储地址为",ptr)
	fmt.Println("main函数中指针的地址为",&ptr)
	fmt.Println("变量val值为",val)
}
/* 值传递 */
func valueTransmit(val int){
	fmt.Println("变量val地址为",&val)
	val=2
}
/* 指针传递 */
func pointerTransmit(ptr *int){
	fmt.Println("指针存储地址为",ptr)
	fmt.Println("指针的地址为",&ptr)
	*ptr=2
}