package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str string="ypf享受编程!"
	fmt.Println(len(str))//长度为16
	/* 转为字节数组 */
	for _,b:=range []byte(str){
		fmt.Printf("%d ",b)//121 112 102 228 186 171 229 143 151 231 188 150 231 168 139 33
	}
	fmt.Println()
	for i,ch:=range str{//ch 就是rune类型 即字符类型
		fmt.Printf("[%d %d]",i,ch)//[0 121][1 112][2 102][3 20139][6 21463][9 32534][12 31243][15 33]
	}
	/* 转为字符数组 */
	fmt.Println()
	for i,ch:=range []rune(str){//ch 就是rune类型 即字符类型
		fmt.Printf("[%d %d]",i,ch)//[0 121][1 112][2 102][3 20139][4 21463][5 32534][6 31243][7 33]
	}
	fmt.Println()
	fmt.Println(utf8.RuneCountInString(str))//8
	bytes :=[]byte(str)
	/* 字节数组转字符 */
	for len(bytes)>0 {
		ch,len :=utf8.DecodeRune(bytes)
		bytes=bytes[len:]
		fmt.Printf("ch=%c ",ch)
	}
}
