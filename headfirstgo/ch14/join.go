package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(JoinWhtiCommas([]string{"apple", "orange", "pear", "banana"}))
}

// JoinWhtiCommas 拼接字符串
func JoinWhtiCommas(phrases []string) string {
	if len(phrases) == 0 {
		return ""
	} else if len(phrases) == 1 {
		return phrases[0]
	} else if len(phrases) == 2 {
		return strings.Join(phrases, " and ")
	} else {
		result := strings.Join(phrases[:len(phrases)-1], ", ")
		result += ", and "
		result += phrases[len(phrases)-1]
		return result
	}
}
