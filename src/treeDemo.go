package main

import (
	"fmt"
	"go-study/tree"
)

func main() {
	root :=tree.Node{Val: 1}
	node2 :=tree.Node{Val: 2}
	node3 :=tree.Node{Val: 3}
	node4 :=tree.Node{Val: 4}
	node5 :=tree.Node{Val: 5}
	root.Left=&node2
	root.Right=&node3
	node2.Left=&node4
	node2.Right=&node5
	//root.Traver()
	/* 中序遍历 同时通过函数编程 统计节点数*/
	count :=0
	root.TraverFunc(func(node *tree.Node) {
		count++
	})
	fmt.Println(count)
	/* 中序遍历 同时通过channel 统计最大值*/
	fmt.Println("-------------")
	max :=0
	c :=root.TraverChan()
	for node :=range c {
		if node.Val>max{
			max=node.Val
		}
	}
	fmt.Println(max)
}
