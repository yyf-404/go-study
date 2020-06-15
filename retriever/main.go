package main

import (
	"fmt"
	"go-study/retriever/mock"
	rea "go-study/retriever/real"
	"go-study/retriever/rp"
)

type Retriever interface {
	Get(url string) string
}
func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}
func main() {
	var retriever Retriever
	retriever=mock.Retriever{"this is a fake website"}
     fmt.Println(download(retriever))//this is a fake website
	fmt.Printf("%T %v\n",retriever,retriever)//mock.Retriever {this is a fake website}
	fmt.Println(retriever.(mock.Retriever))//{this is a fake website}
	//fmt.Println(retriever.(rea.Retriever))panic: interface conversion: main.Retriever is mock.Retriever, not rea.Retriever
	inspect(retriever)
	retriever=&rea.Retriever{}
	fmt.Printf("%T %v\n",retriever,retriever)//*rea.Retriever &{ 0s}
	//fmt.Println(download(retriever))
    var retrieverPoster rp.RetrieverPoster
	retrieverPoster=&rp.RPCometrue{"this is a fake website"}
	fmt.Println(retrieverPoster.Get("http://www.baidu.com"))//this is a fake website
	fmt.Println(rp.Session(retrieverPoster))//another faked  website
}
func inspect(r Retriever) {
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *rea.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
