// 从输入中获取分数然后输出
package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {
	fmt.Print("Enter a Grade: ")
	grade, err := keyboard.GetFloat()
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
