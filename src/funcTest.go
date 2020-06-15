package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("max=",max(1,2))
	fmt.Println(sort(3,4))
	/* 多个返回值，如何只取一个  */
	a,_ :=sort(5,6)
	fmt.Println(a)
	fmt.Println(changeargs(1,2,3))
}
func  max (a int ,b int) int  {
	val:=math.Max(float64(a),float64(b))
	return int(val)
}
func  sort (a int ,b int) (int,int) {
	if a>=b {
		return a,b
	}
		return b,a
}
/* 函数可变参数 */
func changeargs( nums ... int) int{
	sum:=0
	for _,num:= range nums{
		sum+=num
	}
	return sum
}