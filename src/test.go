package main

import "fmt"

const (
	i = iota
	j = iota
	x = iota
)
const xx = iota
const yy = iota
func main(){
	println(i, j, x, xx, yy)
	enums()

}
func enums() {
	const (
		java   = 1
		goland = 2
		cpp    = 3
	)
	fmt.Println(java,goland,cpp)
}