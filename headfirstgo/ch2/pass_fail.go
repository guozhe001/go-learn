// 从输入中获取分数然后输出
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a Grade: ")
	// 从标准输入中打开一个Reader
	reader := bufio.NewReader(os.Stdin)
	// reader读取用户输入的数据，直到出现换行符\n为止
	input, err := reader.ReadString('\n')
	// 如果err改为_，则会reader.ReadString('\n')的错误，继续往下执行，这样会出现非期待的结果
	if err != nil {
		// 打印错误信息并停止程序
		log.Fatal(err)
	}
	// trim字符串，包括换行符、制表符、空格
	input = strings.TrimSpace(input)
	// 将字符串转换成浮点数，第一个参数是字符串，第二个参数是精度
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade == 100 {
		status = "Perfect!"
	} else if grade >= 60 {
		status = "pass!"
	} else {
		status = "fail!"
	}
	fmt.Println("the grade:", grade, "is", status)

}
