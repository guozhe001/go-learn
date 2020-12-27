package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers, err := GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(numbers)
}

// GetFloats 给定一个文件名，读取文件中的数字并以数组的形式反会这些数字
func GetFloats(fileName string) ([3]float64, error) {
	var numbers [3]float64
	// 打开data.txt文件进行读取
	file, err := os.Open(fileName)
	// 如果打开文件出错则输出错误信息并退出程序
	if err != nil {
		return numbers, err
	}
	// 为文件创建一个新的扫描器
	scaner := bufio.NewScanner(file)
	index := 0
	// 逐行读取文件直到文件尾（或出现错误）
	for scaner.Scan() { //从文件中读取一行
		// fmt.Println(scaner.Text()) //打印该行
		// 注意下面的语句也不会报错，下面打印的是值所在的内存地址，上面打印的是值
		// fmt.Println(scaner.Text)
		numbers[index], err = strconv.ParseFloat(scaner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		index++
	}
	// 关闭文件释放资源
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	// 如果扫描文件出现错误，报告错误并退出
	if scaner.Err() != nil {
		return numbers, err
	}
	return numbers, nil
}
