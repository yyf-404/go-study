package main

import (
	"fmt"
	"math/rand"
	"time"
)
func randonVal() chan int{
	out :=make(chan int)
	go func() {
		i:=0
		for  {
			time.Sleep(time.Duration(rand.Intn(2000))*time.Millisecond)
			out <-i
			i++
		}
	}()
	return out
}

func main() {
	var c1,c2 chan int
	c1= randonVal()
	c2=randonVal()
	for  {
		/* 通过select实现非阻塞接收*/
		select {
		case n :=<-c1: fmt.Println(n)
		case n :=<-c2: fmt.Println(n)
		}
	}

}
