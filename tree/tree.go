package tree

import "fmt"

type Node struct {
	Val int
	Left *Node
	Right *Node
}
/* 中序遍历*/
func (node *Node)Traver(){
	if node ==nil {
		return
	}
	node.Left.Traver()
	fmt.Printf("node va; is %d\n",node.Val)
	node.Right.Traver()
}
/* 中序遍历 同时通过函数编程 进行一些操作*/
func (node *Node)TraverFunc(f func(node *Node)){
	if node ==nil {
		return
	}
	node.Left.TraverFunc(f)
	f(node)
	fmt.Printf("node va; is %d\n",node.Val)
	node.Right.TraverFunc(f)
}
/* 中序遍历 同时通过channel 进行一些操作*/
func (node *Node)TraverChan() chan *Node{
    out :=make(chan *Node)
	go func() {
		node.TraverFunc(func(node *Node) {
			out <-node
		})
		close(out)
	}()
     return out
}