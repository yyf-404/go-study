package life

import "fmt"

type life interface {
	Breath()
}
type Human struct {
   Name string
}
func (human Human) Breath(){
	fmt.Println("human can breathing")
}

