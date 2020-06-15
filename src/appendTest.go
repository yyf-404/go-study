package main

import "fmt"

func main() {
	var numbers []int
	/* 允许追加空切片 */
	/* 向切片添加一个元素 */
	numbers = append(numbers, 0)
	printSlice2(numbers)
	/* 同时添加多个元素 */
	numbers =append(numbers,2,3,4)
	printSlice2(numbers)
	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers2:= make([]int,len(numbers),cap(numbers)*2)
	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers2,numbers)
	printSlice2(numbers2)
	numbers3:=numbers[1:3]
	printSlice2(numbers3)
	numbers4:=numbers[0:2]
	printSlice2(numbers4)
}
func printSlice2(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}