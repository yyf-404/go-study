package main

import (
	"fmt"
	sort2 "sort"
)

func main() {
	var arr []int
	arr=make([]int ,3)
	arr[0]=1
	fmt.Println(arr)
    slice1:=make([]int ,3,10)
	fmt.Println(slice1)
	s :=[] int {1,2,3 }
	fmt.Println(s)
	fmt.Printf("len=%d cap=%d slice1=%v\n",len(slice1),cap(slice1),slice1)
	var slice2 []int
	fmt.Printf("len=%d cap=%d slice2=%v isNil=%v\n",len(slice2),cap(slice2),slice2,slice2==nil)
	numbers := []int{0,1,2,3,4,5,6,7,8}
	/* 打印原始切片 */
	printSlice(numbers)
	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])
	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])
	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])
	/* 通过切片截取 创建切片*/
	numbers2 :=numbers[1:6]
	printSlice(numbers2)
	s2 :=[] int {1,3,4,3 }
	sort2.Ints(s2)
	fmt.Println(s2)
}
func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}