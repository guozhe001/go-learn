// 《headfistgo》 第198页的代码贴
package main

import "fmt"

func main() {
	fmt.Println(sum(9, 7))
	fmt.Println(sum(1, 2, 4))
}

func sum(numbers ...int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}
