package main

import (
	"testing"
)

func TestTriangle(t *testing.T) {//测试列表
	tests :=[]struct{//结构体中可以定义各种类型的属性
		a int
		b int
		c int}{
		{3,4,5},
		{5,12,13},
		{8,15,17},
		{30000,40000,50000},
	}
	for _,tt :=range tests{
		if actual :=triangle(tt.a,tt.b) ;actual!=tt.c{//测试所有结果
		t.Errorf("triangle(%d, %d); "+
				"got %d; expected %d",
				tt.a, tt.b, actual, tt.c)
		}
	}
}
 func BenchmarkTest(b *testing.B){
	val1,val2:=30000,40000
	ans :=50000
	for i:=0;i<b.N;i++{
		actual :=triangle(val1,val2)
		if actual !=ans{
			b.Errorf("triangle(%d %d)  actual is %d; "+
				"expected %d",
				val1,val2,actual,  ans)

		}
	}
 }