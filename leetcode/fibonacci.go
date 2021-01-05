package main

import "fmt"

const n = 20

func main() {
	var numbers [n + 1]int
	for i := 0; i <= n; i++ {
		if i == 0 {
			numbers[i] = 0
		} else if i == 1 {
			numbers[i] = 1
		} else {
			numbers[i] = numbers[i-1] + numbers[i-2]
		}
		fmt.Print(numbers[i], " ")
	}
	fmt.Println()
}
