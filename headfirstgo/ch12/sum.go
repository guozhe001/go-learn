package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// OpenFile 打开文件
func OpenFile(name string) (*os.File, error) {
	fmt.Println("Opening", name)
	return os.Open(name)
}

// CloseFlie 关闭文件
func CloseFlie(file *os.File) {
	fmt.Println("Closing file")
	file.Close()
}

// GetFloats 从文件中获取float64切片
func GetFloats(name string) ([]float64, error) {
	file, err := OpenFile(name)
	if err != nil {
		return nil, err
	}
	defer CloseFlie(file)
	scanner := bufio.NewScanner(file)
	var numbers []float64
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}

func main() {
	if err := Socilize(); err != nil {
		log.Fatal(err)
	}
	numbers, err := GetFloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	fmt.Printf("Sum: %0.2f\n", sum)
}

// Socilize 测试顺序
func Socilize() error {
	defer fmt.Println("Goodbye!")
	fmt.Println("Hello")
	// 如果出现两个以上defer，则按照倒序执行
	defer fmt.Println("hi")
	return fmt.Errorf("i don't want do talk")
	// 在return之后的defer语句也不会被执行
	defer fmt.Println("haha")
	fmt.Println("nice Weather")
	return nil
}
