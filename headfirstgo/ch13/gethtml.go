package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// go a()
	// go b()
	// time.Sleep(time.Second)
	// fmt.Println()
	// fmt.Println("end main")
	// abcdef()
	// receiveAndSend()
	urls := []string{"http://www.example.com", "http://www.baidu.com", "http://www.sina.com.cn/"}
	channel := make(chan urlSize)
	for _, url := range urls {
		go responseSize(url, channel)
	}
	for i := 0; i < len(urls); i++ {
		page := <-channel
		fmt.Printf("%s: %d\n", page.url, page.size)
	}
}

type urlSize struct {
	url  string
	size int
}

func responseSize(url string, channel chan urlSize) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// channel <- urlSize{url, len(body)}
	channel <- urlSize{url: url, size: len(body)}
}

func a() {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
}

func b() {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
}

func abcdef() {
	channel := make(chan string)
	channel2 := make(chan string)
	go abc(channel)
	go def(channel2)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel2)
	fmt.Println(<-channel2)
	fmt.Println(<-channel2)
}

func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func receiveAndSend() {
	myChannel := make(chan string)
	// 刚开始运行时 receiving和sending两个goroutine都在休眠

	go send(myChannel)
	reportNap("receiving goroutine", 5)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}

func send(channel chan string) {
	reportNap("sending goroutine", 2)
	fmt.Println("sending value***")
	channel <- "a"
	fmt.Println("sending value***")
	channel <- "b"
}

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
		time.Sleep(time.Second)
	}
	fmt.Println(name, "wakes up!")
}
