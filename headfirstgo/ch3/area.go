// 第三章，方法定义章节练习
package main

import "fmt"

func main() {
	var total float64
	total += paintNeeded(2.4, 5.6)
	total += paintNeeded(2.6, 9.6)
	fmt.Printf("total is %.2f\n", total)
}

func paintNeeded(weidth float64, higth float64) float64 {
	area := (weidth * higth) / 10
	fmt.Printf("%.2f liters needed\n", area)
	return area
}
