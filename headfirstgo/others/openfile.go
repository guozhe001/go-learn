package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const fileName string = "a.txt"

func main() {
	write(fileName)
	read(fileName)
}

func write(name string) {
	// Go允许你在程序代码中使用八进制表示法来编写数字，任何以0开头的数字序列都将被视为八进制数
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.FileMode(0600))
	check(err)
	_, err = file.Write([]byte("amazing!\n"))
	check(err)
	err = file.Close()
	check(err)
}

func read(name string) {
	file, err := os.OpenFile(name, os.O_RDONLY, os.FileMode(0600))
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	check(scanner.Err())
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
