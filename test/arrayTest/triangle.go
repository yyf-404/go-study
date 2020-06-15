package main

import "math"
/* 勾股定理计算 */
func triangle(a int,b int) int {
	tem :=math.Pow(float64(a),float64(2))+math.Pow(float64(b),float64(2))
	return int(math.Sqrt(tem))
}
