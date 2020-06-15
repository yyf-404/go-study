package main

import (
	"fmt"

)

func main() {
	var map1 map[string] int
	map1=make(map[string]int)
	map1["a"]=1
	map1["b"]=2
	map1["c"]=3
	/* 读取单个元素 */
	fmt.Println(map1["a"])
	/* 遍历操作 */
	for key,value:= range map1{
    fmt.Println(key+"=",value)
	}
	/*查看元素在集合中是否存在 */
	number,flag :=map1["c"]
	isHas(flag,number)
    delete(map1,"c")
	number2,flag2 :=map1["c"]
	isHas(flag2,number2)

}
func isHas(flag bool, number int) {
	if(flag) {
		fmt.Println("c字母对应数字存在为",number)
	}else{
		fmt.Println("c字母对应数字不存在")
	}
}
