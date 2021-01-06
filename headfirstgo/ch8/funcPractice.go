package main

import "fmt"

type MyType string

func main() {
	a := MyType("a")
	a.sayHi()

	b := MyType("b")
	b.sayHi()

	two := Number(2)
	fmt.Println("Origin value of number:", two)
	two.Double()
	fmt.Println("number after calling Double:", two)
	two.double()
	fmt.Println("number after calling double:", two)
	// 指针类型调用非指针类型的接收器的方法，也会被Go自动获取指针的值然后传递给方法。
	pointerTwo := &two
	pointerTwo.Double()
	fmt.Println("number after calling Double:", two)

	// 下面是第282页实践出真知
	number := Number(4)
	number.double()
	number.Display()

	// 在一个没有保存到变量的值上直接调用一个要求接收器的指针的方法
	// Number(5).double()
	// 上面的一行会报如下错误：
	// ./funcPractice.go:29:11: cannot call pointer method on Number(5)
	// ./funcPractice.go:29:11: cannot take the address of Number(5)

	// 在一个没有保存到变量的值上直接调用一个要求接收器的值的方法时是不会报错的
	Number(5).Double()
}

func (m MyType) sayHi() {
	fmt.Println("Hi from", m)
}

type Number int

func (n Number) Double() {
	n *= 2
}

func (n *Number) double() {
	*n *= 2
}

func (n *Number) Display() {
	fmt.Println(*n)
}
