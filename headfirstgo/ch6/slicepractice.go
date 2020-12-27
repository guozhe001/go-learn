// Package main 切片slice练习
package main

import (
	"fmt"
)

func main() {
	fmt.Println("parctice1========================")
	parctice1()
	fmt.Println("parctice2========================")
	mySlice := practice2()
	fmt.Println("parctice3========================")
	practice3()
	fmt.Println("parctice4========================")
	practice4()
	fmt.Println("for loop print mySlice========================")
	printSlice(mySlice)
	fmt.Println("sliceIsArrayView========================")
	sliceIsArrayView()
}

func printSlice(mySlice []string) {
	// for range
	for _, v := range mySlice {
		fmt.Println(v)
	}

	// for index
	for i := 0; i < len(mySlice); i++ {
		fmt.Println("i=", i, "value=", mySlice[i])
	}
}

func parctice1() {
	// 此时只是定义了一个名为mySlice的slice类型的变量，但是并没有自动创建一个切片
	var mySlice []string
	// 调用make方法之后才可以操作mySlice变量，否则会报错如下：
	// panic: runtime error: index out of range [0] with length 0
	mySlice = make([]string, 1)
	// 如果make函数的第二个值是0，则给index为0赋值时也会报index out of range
	mySlice[0] = "do"
	fmt.Println(mySlice[0])
}

func practice2() []string {
	// 短变量声明方式
	mySlice := make([]string, 7)
	mySlice[0] = "do"
	mySlice[1] = "re"
	mySlice[2] = "mi"
	fmt.Println(mySlice[2])
	return mySlice
}

func practice3() {
	// slice的第三种声明方式
	// 与数组的区别是方括号内没有指定长度，仅此而已
	mySlice := []string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	// 下面的赋值方式会报错如下：
	// panic: runtime error: index out of range [7] with length 7
	// 使用下标的方式操作mySlice与数组没有任何差别
	// mySlice[len(mySlice)] = "hello"
}

func practice4() {
	// 切片是数组的一部分，（这一部分也可以是全部）
	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println(primes)
	// 使用数组创建切片，左闭右开
	aSlice := primes[1:4]
	fmt.Println(aSlice)
	// 如果创建切片的值超出了数组的范围，直接报编译错误如下：
	// index 6 (constant of type int) is out of boundscompiler
	// bSlice := primes[0:6]
	// fmt.Println(bSlice)

	// 可以忽略开始和结束的索引
	// 忽略开始索引，表示开始索引为0。下面的定义与primes[0:4]等效
	bSlice := primes[:4]
	fmt.Println(bSlice)
	// 忽略结束索引，表示直到最后一个元素为。下面的定义与primes[1:len(primes)]等效
	cSlice := primes[1:]
	fmt.Println(cSlice)
}

// sliceIsArrayView *注意：由于切片只是底层数组内容的视图，如果你修改底层数组，这些变化也会反映到切片*
func sliceIsArrayView() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(notes)
	aSlice := notes[0:5]
	fmt.Println(aSlice)
	// 修改原始的数组的第一个元素
	notes[0] = "a"
	// slice的第一个元素也已经变成了"a"
	fmt.Println(aSlice[0])
	// 操作切片也修改了切片对应的底层的数组
	aSlice[0] = "do"
	fmt.Println(notes[0])
}
