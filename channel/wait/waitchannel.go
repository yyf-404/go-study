package main

import (
	"fmt"
)

type  worker struct {
	 c chan int
	 done chan bool //标记是否完成
}
func doWorker(id int,c chan int,done chan bool) {
	for {
		n := <-c
		fmt.Printf("worker %d receive %c\n",id,n)
		//如果需要返回多次done<-true标记结束，并且一直在传入数据，会产生死锁
		//所以要重开一个 goroutine 返回消息
		go func() {done<-true}()
	}
}
func channelWait() {
	var workers [10]worker

	for i:=0;i<10;i++{
		workers[i]=worker{c: make(chan int),done: make(chan bool)}
		go doWorker(i,workers[i].c,workers[i].done)

	}
	for i,w :=range workers{
		w.c <- 'a'+i

	}
	for i,w :=range workers{
		w.c <- 'A'+i
	}
	//接收消息 保证工作线程执行完
	for _,w:=range workers{
		<-w.done
		<-w.done
	}
}
func main() {
	channelWait()
}
