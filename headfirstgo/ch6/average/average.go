package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args)
	// 第一个元素是程序本身
	fmt.Println(os.Args[1:])
	// 调用可变长参数函数时，可以传入切片，切片后面加...即可
	average(convert(os.Args[1:]...)...)
}

func convert(args ...string) []float64 {
	var numbers []float64
	for _, arg := range args {
		number, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func average(numbers ...float64) float64 {
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}
