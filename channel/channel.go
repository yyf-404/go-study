package main

import (
	"fmt"
	"time"
)

func worker(c chan int) {
	for{
		n := <-c
		fmt.Printf("worker receive %c\n",n)
	}
}
func createWorker() chan<- int{
	c :=make(chan int)
   go worker(c)
	return c
}
func channelDemo() {
	c := make(chan int)
	go worker(c)
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}
func channelDemo2() {
	var c chan <-int
	c =createWorker()
	c <- 1
	c <- 2
	//n :=<-c  无法编译  chan <-int只能发数据
	time.Sleep(time.Millisecond)
}
/* 缓存channel */
func bufferChannel(){
	c := make(chan int,3)//定义缓冲区大小为3
	go worker(c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}
func closeChannel(){
	c := make(chan int)
	go worker2(c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	/* 关闭channel */
	close(c)
	time.Sleep(time.Millisecond)
}
func worker2(c chan int) {
	for{
		n,ok := <-c
		if !ok{ //要判断是否发送完毕 不然会一直传送空串
			break
		}
		fmt.Printf("worker receive %c\n",n)
	}
}
func main() {
//channelDemo()
	//channelDemo2()
	//bufferChannel()
      closeChannel()
}
