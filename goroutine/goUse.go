package main

import (
	"fmt"
	"time"
)

func main() {
	var arr [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			//fmt.Println(i)//io阻塞的方式可以使协程交出控制权
			for {
				arr[i]++
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(arr)
}
