package main

import "fmt"

func add() func(int) int {
	sum := 0
	return func(value int) int {
		sum += value
		return sum
	}
}
/* 正统函数式编程 （不使用状态变量）*/
type iAdder func(int) (int,iAdder) //定义接口
func add2(base int) iAdder {
	return func(value int) (int,iAdder) {
		return base+value,add2(base+value)
	}
}
func main() {
	a := add()
	for i := 1; i < 10; i++ {
		fmt.Printf("0+1+...+%d =%d\n", i, a(i))
	}
	a2 := add2(0)
	var sum int
	for i := 1; i < 10; i++ {
		sum,a2 =a2(i)
		fmt.Printf("0+1+...+%d =%d\n", i, sum)
	}
}
