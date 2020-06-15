package main

import "fmt"

func main() {
	var arr [3]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3, 4}
	arrPrint(arr)
	arrPrint(arr2)
	//arrPrint(arr3) 无法通过编译
	fmt.Println(arr3)
	fmt.Println(arr)
	arr4 := []int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr4[2:6]
	s2 := s1[3:5]
	//s3:=s1[3:7] 运行时报错 向后扩展无法超过cap大小
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)
	fmt.Println("--------------")
	s4 := append(s1, 10)
	s5 := append(s4, 11)
	s6 := append(s5, 12)
	fmt.Println("s4,s5,s6 =", s4, s5, s6)
	fmt.Println(arr4)
	/* 删除中间元素 */
	s6 = append(s6[:3], s6[4:]...)
	fmt.Println(s6) //输出[2 3 4 10 11 12]
	/* 删除首元素 */
	s6 = s6[1:]
	fmt.Println(s6) //[3 4 11 12]
	/* 删除尾元素 */
	s6 = s6[:len(s6)-1]
	fmt.Println(s6)//[3 4 10 11]

}
func arrPrint(arr [3]int ){
	fmt.Println(arr)
	arr[0]=-1
}
