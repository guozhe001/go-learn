package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		// 读取文件内容
		data, err := ioutil.ReadFile(filename)
		fmt.Println(reflect.TypeOf(data))
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		// 把读取到的文件数据按照换行符\n进行分割，并循环
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
