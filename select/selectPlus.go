package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randonVal2() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}
func worker2(c chan int) {
	for {
		time.Sleep(2000 * time.Millisecond)
		n := <-c
		fmt.Printf("worker2 receive %d\n", n)
	}
}
func createworker2() chan int {
	c := make(chan int)
	go worker2(c)
	return c
}
func main() {
	var c1, c2 chan int
	c1 = randonVal2()
	c2 = randonVal2()
	var values []int
	n := 0
	worker2 := createworker2()
	/* 通过返回一个数据给 channel tm 标记计时结束*/
	tm := time.After(10 * time.Second)
	/* 定时器：定时返回数据给channel*/
	tick := time.Tick(time.Second)
	for {
		/* 通过select实现非阻塞接收*/
		/* 通过 select 不会发送数据到 nil channel 来实现有数据才进行发送*/
		var w chan int
		var activeVal int
		if len(values) > 0 {
			w = worker2
			activeVal = values[0]
		}
		select {
		case n = <-c1:
			values = append(values, n)
		case n = <-c2:
			values = append(values, n)
			//w为nil时 阻塞
		case w <- activeVal:
			values = values[1:]
		/* 如果800毫秒 没有获得数据 就输出超时*/
		case <-time.After(800 * time.Millisecond):
			fmt.Println("time out...")
		case <-tick:
			fmt.Println("定时输出队列长度 %d", len(values))
		case <-tm:
			fmt.Println("system exit.....")
			return
		}
	}
}
