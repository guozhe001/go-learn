// 使用Join输出变量名
package main

import (
	"fmt"
	"os"
	"strings"
)

func main () {
	fmt.Println(os.Args[0])
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
	practice12()
}

func practice12() {
	for index,value := range os.Args[:] {
		fmt.Println(index, value)
	}
}