// Package main 数组练习
package main

import (
	"fmt"
	"time"
)

// printStringArray 打印字符串类型的数组
func printStringArray() {
	fmt.Println("printStringArray=================================")
	// 定义数组的方法1
	var strs [7]string
	strs[0] = "do"
	strs[1] = "re"
	strs[2] = "mi"
	fmt.Println(strs[0])
	fmt.Println(strs[1])

	// 定义数组的方法2:
	var notes [7]string = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	// 打印数组的值: [do re mi fa so la ti]
	fmt.Println(notes)
	// 打印数组的定义: [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Printf("%#v\n", notes)

	// 定义数组的方法3:
	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println(primes)
	fmt.Printf("%#v\n", primes)
}

func printIntArray() {
	fmt.Println("printIntArray=================================")
	var ints [3]int
	ints[0] = 1
	fmt.Println(ints[0])
}

func printTimeArray() {
	fmt.Println("printTimeArray=================================")
	var times [3]time.Time
	times[0] = time.Now()
	fmt.Println(times[0])
}

func forLoopArray() {
	fmt.Println("forLoopArray=================================")
	// 遍历一个数组，方式1：
	fmt.Println("遍历一个数组，方式1：for i := 0; i < len(notes); i++ {=================================")
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	for i := 0; i < len(notes); i++ {
		fmt.Println(notes[i])
	}
	// 推荐使用下面两种方式
	// 遍历一个数组，方式2：
	fmt.Println("遍历一个数组，方式2：for index, value := range notes {=================================")
	for index, value := range notes {
		fmt.Println(index, value)
	}
	// 遍历一个数组，方式2.1：
	fmt.Println("遍历一个数组，方式2.1：for _, value := range notes  {=================================")
	for _, value := range notes {
		fmt.Println(value)
	}

}

func main() {
	printStringArray()
	printIntArray()
	printTimeArray()
	forLoopArray()
}
