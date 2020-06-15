package main

import "fmt"

func main() {
	var a int =10 //声明实际变量
	var p *int    //声明指针变量
	p=&a         //将变量的存储地址赋值给指针
	fmt.Printf("a 变量的地址是: %x\n", &a  )
	/* 指针变量的存储地址 */
	fmt.Printf("p 变量储存的指针地址: %x\n", p )
	/* 使用指针访问值 */
	fmt.Printf("*p 变量的值: %d\n", *p )
	var  ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr )
	/* 空指针判断 */
	fmt.Println( ptr==nil )
	pointer()
	val1:=1
	val2:=2
	fmt.Printf("交换前 val1 的值 : %d\n", val1)
	fmt.Printf("交换前 val2 的值 : %d\n", val2 )
	swap(&val1,&val2)
	fmt.Printf("交换后 val1 的值 : %d\n", val1)
	fmt.Printf("交换后 val2 的值 : %d\n", val2 )
}
func pointer(){
	var a = 5
	var ptr *int
	var pptr **int
	ptr=&a
	pptr=&ptr
	fmt.Printf("变量 a = %d\n", a )
	fmt.Printf("变量 a地址 = %d\n", &a )
	fmt.Println("------------------")
	fmt.Printf("指针变量 *ptr = %d\n", *ptr )
	fmt.Printf("指针变量指向地址值 ptr = %d\n", ptr )
	fmt.Printf("指针变量地址 ptr = %d\n",& ptr )
	fmt.Println("------------------")
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
	fmt.Printf("指向指针的指针变量指向地址值 pptr = %d\n", pptr)
	fmt.Printf("指向指针的指针变量地址 pptr = %d\n", &pptr)
}
func swap(a *int,b *int){
	 temp:=*a;
	 *a=*b
	 *b=temp
}