package main

import (
	"fmt"
	"io/ioutil"
)
func main() {
	a:=10
	if a>=10{
		fmt.Println("yes")
	}else if a<0{
		fmt.Println("no")
	}else {
		fmt.Println("no")
	}
	for  i:=0;i<10;i++ {
		fmt.Println(i)
	}
	var arr[] int=[]int{1,2}

	for i,s:=range arr {
		fmt.Println(i,s)
	}
	var arr2[10] int
	for i:=0;i<10;i++ {
		arr2[i]=i;
		fmt.Println("arr[",i,"]=",arr2[i])
	}
	ifFunc()
	fmt.Println(switchFunc(90))
	fmt.Println(switchFunc(80))
	fmt.Println(switchFunc(101))
}
func ifFunc(){
	const filename="abc.txt"
	if contents,err:=ioutil.ReadFile(filename);err==nil{
		fmt.Printf("%s\n",contents)
	}else {
		fmt.Println("some err has happened",err)
	}
}
func switchFunc(score int) string {
	res := ""
	switch {
	case score > 100 || score < 0:
		res = "score input is err"
	case score >= 90:
		res = "A"
	case score >= 60:
		res = "B"
	case score <= 40:
		res = "C"
	}
	return res
}