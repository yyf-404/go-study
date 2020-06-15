package main

import (
	"fmt"
	"go-study/pkg"
)
/* 组合*/
type SuperMan struct {
	human *life.Human
}

func (superMan SuperMan)fly()  {
	fmt.Println("human can fly....")
}
func main() {
	var superMan SuperMan = SuperMan{human: &life.Human{"zzz"}}
	superMan.human.Breath()
	superMan.fly()
	var queue life.Queue
	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Pop())
	fmt.Println(queue.IsEmput())
	fmt.Println(queue.Pop())
	fmt.Println(queue.IsEmput())
	queue.Push("aaa")
	fmt.Println(queue.Pop())
}
