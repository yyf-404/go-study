package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fibonacci() func() int{ //斐波那契数列
	a,b:=0,1
	return func() int {
		a,b=b,a+b
		return a
	}
}
type intGen func() int
func (g intGen) Read( //函数类型也可以实现接口
	p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d ", next)
	return strings.NewReader(s).Read(p)
}
func printFileContents(reader io.Reader) { //定义reader内容的读取方法
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}


func main() {
	f:=fibonacci()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	var f2 intGen = fibonacci()
	printFileContents(f2)
}
