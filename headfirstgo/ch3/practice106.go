// 第106页的练习题答案，指针
package main

import "fmt"

func main() {
	var myInt int
	var myIntPointor *int
	myInt = 42
	myIntPointor = &myInt
	fmt.Println(*myIntPointor)
}

func double(number *int) {
	*number *= 2
}
