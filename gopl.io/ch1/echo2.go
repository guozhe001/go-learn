// echo2输出其命令行参数
package main

import (
	"fmt"
	"os"
)

func main() {
    s, sep := "", ""
	for _, arg := range os.Args[1: ] {
		s += sep + arg;
		sep = " "
	}
	fmt.Println(s)
}

// 下面的声明字符串的方式是等价的
// s := ""
// var s string
// var s string = ""
// var s = ""
