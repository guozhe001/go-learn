// Package arraypractice 数组练习
package arraypractice

import "fmt"

// PrintStringArray 打印字符串类型的数组
func PrintStringArray() {
	// strs := [7]string{"do", "re", "mi"}
	var strs [7]string
	strs[0] = "do"
	strs[1] = "re"
	strs[2] = "mi"
	fmt.Println(strs[0])
	fmt.Println(strs[1])
}
