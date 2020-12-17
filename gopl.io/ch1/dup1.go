// dup1 输出标准输入中出现次数大于1的行，前面是次数
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	mapPractice()
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
		/*
		 * 上面的语句等价于
		 * line := input.Text()
		 * counts[line] = counts[line] + 1
		 */
		// 注意：忽略input.Err()中可能的错误
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", line, n)
			}
		}
	}
}

/**
 * map练习:
 * map里面的建的迭代顺序不是固定的，通常是随机的，每次运行都不一致。
 * 这是有意设计的，以防止程序依赖某种特定的序列，此处不对排序做任何保证。
 */
func mapPractice() {
	helloMap := make(map[string]int)
	helloMap["hello"] = 1
	helloMap["h"] = 3
	helloMap["he"] = 2
	for key, value := range helloMap {
		fmt.Println(key, value)
	}
}