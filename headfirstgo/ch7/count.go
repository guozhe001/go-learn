package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	lines, err := readStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
	fmt.Println(solution1(lines), "is success!")
	fmt.Println(solution2(lines), "is success!")
	mapPractice()
	mapPractice2()
	mapForLoop()
}

func solution2(lines []string) string {
	fmt.Println("solution2=============================")
	nameMap := make(map[string]int)
	for _, name := range lines {
		nameMap[name]++
	}

	maxCount := 0
	var maxName string
	for key, value := range nameMap {
		if value > maxCount {
			maxCount = value
			maxName = key
		}
		fmt.Println(key, value)
	}
	return maxName
}

func mapForLoop() {
	fmt.Println("mapForLoop=============================")
	counts := map[string]int{"a": 1, "b": 2, "c": 3}
	// 使用key，value接收range map的返回
	for key, value := range counts {
		fmt.Println("key:", key, "value:", value)
		// 上面的输出结果情况之一如下： 说明map在遍历的时候是随机的，并不是有序的
		// key: c value: 3
		// key: a value: 1
		// key: b value: 2
	}

	// 如果你想要按照一定的顺序“遍历map”，则需要额外的写代码：
	var names []string
	for key := range counts {
		names = append(names, key)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Println("key:", name, "value:", counts[name])
	}

	// 只接收key，value可以省略
	for key := range counts {
		fmt.Println("key:", key)
	}

	// 只接收value，但是key需要占位符
	for _, value := range counts {
		fmt.Println("value:", value)
	}
}

// 使用ok来判断是没有赋值还是赋值了但是是“零值”
// b和c都是0，但是b真的赋值是0，c没有赋值默认是“零值”
func mapPractice2() {
	counts := map[string]int{"a": 1, "b": 0}
	a, ok := counts["a"]
	fmt.Println("a", a, ok)
	// 也可以这么写，只接受一个返回结果，就是key对应的值
	a = counts["a"]
	// 或者前面的值忽略，只用于判断是否存在这个key
	fmt.Println("a", a)
	_, ok = counts["a"]
	fmt.Println("a is", ok)
	b, ok := counts["b"]
	fmt.Println("b", b, ok)
	c, ok := counts["c"]
	fmt.Println("c", c, ok)
	// 使用delete内置函数来删除映射的key和对应的value
	fmt.Println("delete key b")
	delete(counts, "b")
	b, ok = counts["b"]
	fmt.Println("b", b, ok)
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
	isPrame1 := map[int]bool{}
	fmt.Println(isPrame1)
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
