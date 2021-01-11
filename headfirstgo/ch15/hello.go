package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	testFuncFeatures()
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/salut", frenchHandler)
	http.HandleFunc("/namaste", hindiHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("hello web!")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello web!")
}

func frenchHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Salut web!")
}

func hindiHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Namaste web!")
}

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func testFuncFeatures() {
	// 下面是一级函数的练习
	var myFunc func(string)
	myFunc = sayHello
	myFunc("Go")
	myFunc2 := sayHi
	myFunc2()
	twice(sayHi)
	var mathFunc func(int, int) float64
	mathFunc = divide
	result := mathFunc(2, 5)
	fmt.Println(result)
}

// sayHi 测试一级函数
func sayHi() {
	fmt.Println("Hi!")
}

// sayHello 测试带参数的一级函数
func sayHello(name string) {
	fmt.Println("hello", name)
}

func twice(theFunc func()) {
	theFunc()
	theFunc()
}

func divide(a int, b int) float64 {
	return float64(a) / float64(b)
}
