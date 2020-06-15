package main

import "fmt"

type Phone interface {
	call()
}
type Huawei struct {
}

func (huawei Huawei) call() {
	fmt.Println("huawei oncall....")
}

type Iphone struct {
}
func (iphone Iphone) call() {
	fmt.Println("iphone oncall....")
}
func main() {
     var phone Phone
     phone=new(Huawei)
     phone.call()
     phone=new(Iphone)
     phone.call()
}
