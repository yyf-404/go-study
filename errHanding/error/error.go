package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	err:=errors.New("this is a new err") //自定义异常
	fmt.Println(err)
	if pathError,ok:=err.(*os.PathError);!ok{
		panic(err) //抛出异常
	}else{
		fmt.Println(pathError)
	}
}
