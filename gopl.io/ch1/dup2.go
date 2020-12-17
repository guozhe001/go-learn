// dup2 打印输入中多次出现的行的个数和文本
// 它从stdin或指定的文件列表中读取
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
		print(counts, nil)
	} else {
		for _, arg := range files {
			/*
			 * os.Open返回两个值：
			 * 第一个是打开的文件（*os.File）
			 * 第二个返回值一个内置的error类型的值。如果err等于特殊的内置nil值，标准文件成功打开；文件再被读到结尾的时候，Close函数关闭文件，然后释放相应的资源
			 * 另一方面，如果err不是nil，说明出错了。这时error的值描述错误原因。
			 *
			 *
			 */
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			countLines(f, counts)
			print(counts, f)
			f.Close()
		}
	}

}

func print(counts map[string]int, f *os.File) {
	haveRepeatLine := false
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", line, n)
			haveRepeatLine = true
		}
	}
	// 练习1.4，输出出现重复行的文件名称
	if haveRepeatLine && f != nil {
		fmt.Println("filename=", f.Name())
	}
}

/**
 * map是一个使用make创建的数据结构的引用。
 * 当一个map被传递给一个函数时，函数接收到这个引用的副本，
 * 所以被调用函数中对于map数据结构中的改变对函数调用者使用的map饮用也是可见的。
 *
 */
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// 注意：input.Err()中可能的错误
}
