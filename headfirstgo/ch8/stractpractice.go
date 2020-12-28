package main

import "fmt"

func main() {
	defineStruct()
}

func defineStruct() {

	var myStruct struct {
		number float64
		word   string
		toggle bool
	}

	fmt.Printf("%#v\n", myStruct)

	myStruct.number = 3.14
	fmt.Println(myStruct.number)
	showInfo(getCar("保时捷", 323))
}

func getCar(name string, topSpeed float64) car {
	var c car
	c.name = name
	c.topSpeed = topSpeed
	return c
}

func showInfo(c car) {
	fmt.Println("Name:", c.name)
	fmt.Println("Top Speed:", c.topSpeed)
}

type subsriber struct {
	name   string
	rate   float64
	active bool
}

type car struct {
	name     string
	topSpeed float64
}
