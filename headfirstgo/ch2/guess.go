package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// 获取一个随机需要猜测的目标数字
	target := getRandom()
	fmt.Println("i have a random number between 1 and 100.")
	fmt.Printf("can you guess is?")
	// 是否猜对
	var success = false
	reader := bufio.NewReader(os.Stdin)
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("you have ", 10-guesses, "guesses left。")
		fmt.Println("Make your guess:")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		// 将字符类型转换成整数类型
		number, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if number < target {
			fmt.Println("哎呀，你猜低了。")
		} else if number > target {
			fmt.Println("哎呀，你猜高了。")
		} else {
			fmt.Println("干得好，你猜对了！")
			success = true
			break
		}
	}
	if !success {
		fmt.Println("对不起，你没有猜对。他是：", target)
	}
}

func getRandom() int {
	rand.Seed(time.Now().Unix())
	target := rand.Intn(100) + 1
	return target
}
