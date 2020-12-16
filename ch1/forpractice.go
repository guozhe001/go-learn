// for循环练习
package main

import "fmt"

func main() {
	for1()
	fmt.Println("=======================")
	for2()
	fmt.Println("=======================")
	for3()
}

// 三个条件都写的for循环
func for1() {
	// 初始化，条件，循环结束之后执行的post
	// for initializtion; condition; post
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}


func for2() {
	var i = 0
    for i <10 {
		fmt.Println(i)
		i ++
	}
}

func for3() {
	var i = 0
	for {
		fmt.Println(i)
		if i==10 {
			break
		}
	    i ++
	}
}
