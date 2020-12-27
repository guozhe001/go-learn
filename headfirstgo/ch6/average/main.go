package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers, err := getFloats()
	if err != nil {
		log.Fatal(err)
	}
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	average := sum / float64(len(numbers))
	fmt.Printf("average=%0.2f\n", average)
	oneInt(1)
	severalInts(1, 23, 34, 12)
	severalInts(1)
	// 可变长参数的函数，可以不传入任何参数
	severalStrings()
	maxinum(1, 2, 3, 4, 61, 34, 23, 545, 78, 89, 45, 76, 72)
	inRange(10, 30, 23, 43, 45, 12, 1, 34, 2, 4, 9, 88)
	nums := []float64{23, 43, 45, 12, 1, 34, 2, 4, 9, 88}
	inRange(10, 30, nums...)
}

func getFloats() ([]float64, error) {
	var numbers []float64
	for _, v := range os.Args[1:] {
		number, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	return numbers, nil
}

func severalInts(num int, nums ...int) {
	fmt.Println("severalInts=========================")
	var sum int
	for _, v := range nums {
		sum += v
	}
	fmt.Println(sum)
}

func oneInt(num int) {
	fmt.Println("oneInt=========================")
	fmt.Println(num)
}

func severalStrings(strs ...string) {
	for _, str := range strs {
		fmt.Println(str)
	}
}

func maxinum(numbers ...int) int {
	fmt.Println("maxinum=================", numbers)
	var max int
	for _, v := range numbers {
		if v > max {
			max = v
		}
	}
	return max
}

func inRange(min float64, max float64, numbers ...float64) []float64 {
	fmt.Println("inRange=======================", min, max, numbers)
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	fmt.Println(result)
	return result
}
