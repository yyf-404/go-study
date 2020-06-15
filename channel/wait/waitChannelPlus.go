package main

import (
	"fmt"
	"sync"
)

type  worker2 struct {
	c chan int
	wg *sync.WaitGroup  //系统提供工具类
}
func doWorker2(id int,c chan int,wg sync.WaitGroup) {
	for {
		n := <-c
		fmt.Printf("worker %d receive %c\n",id,n)
		wg.Done()
	}

}
func channelWait2() {
	var workers [10]worker2
	var wg sync.WaitGroup
	for i:=0;i<10;i++{
		workers[i]=worker2{c: make(chan int),wg: &wg}
		go doWorker2(i,workers[i].c,wg)
	}
	//等待20个信息传送完
    wg.Add(20)
	for i,w :=range workers{
		w.c <- 'a'+i

	}
	for i,w :=range workers{
		w.c <- 'A'+i
	}
	//保证工作线程执行完
	wg.Wait()
}
func main() {
	channelWait2()
}
