package main

import (
	"errors"
	"fmt"
)

func tryRecover(){
	defer func() {
		r:=recover()
		if err,ok := r.(error);ok{
			fmt.Println("Error occurred:",err)//如果是错务直接输出
		}else{
			panic(err)//不是直接抛出
		}
	}()
	panic(errors.New("this is a new error"))
	//panic(123)
}

func main() {
	tryRecover()
}
