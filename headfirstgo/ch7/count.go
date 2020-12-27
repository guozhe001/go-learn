package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	lines, err := readStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
	fmt.Println(solution1(lines), "is success!")
	mapPractice()
}

func mapPractice() {
	fmt.Println("mapPractice=============================")
	// 定义map方式1:先声明，再创建=============================
	// var isPrame map[int]bool
	// isPrame = make(map[int]bool)
	// 定义map方式2:短变量声明方式=============================
	// isPrame := make(map[int]bool)
	// // map赋值
	// isPrame[4] = false
	// isPrame[7] = true
	// 定义map方式3:映射字面量=============================
	isPrame := map[int]bool{4: false, 7: true}
	// map取值
	fmt.Println(isPrame[4])
	// 遍历map
	for key, value := range isPrame {
		fmt.Println("key=", key, ",value=", value)
	}
}

// 选票解决方案1:给定一个切片，返回切片中出现次数最多的字符串
func solution1(lines []string) string {
	fmt.Println("solution1=============================")
	var names []string
	var counts []int
	for _, line := range lines {
		var match bool
		for index, name := range names {
			if name == line {
				match = true
				counts[index]++
			}
		}
		if !match {
			names = append(names, line)
			counts = append(counts, 1)
		}
	}

	var successIndex = 0
	for i, v := range counts {
		if v > counts[successIndex] {
			successIndex = i
		}
	}

	return names[successIndex]
}

func readStrings(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, err
	}
	return lines, nil
}
