package main

import "fmt"

func main() {
	defineStruct()
	test()
	fmt.Println(Gallons(12))
	fmt.Println(Liters(12))
	fmt.Println(Milliliters(12))
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

// 具有底层基础类型的定义类型

type Liters float64
type Gallons float64
type Milliliters float64

func test() {
	carFuel := Gallons(1.2)
	busFuel := Liters(2.5)
	// carFuel += Liters(8.0)
	// busFuel += Gallons(30.0)
	// 下面两行运行时报错如下：
	// ./stractpractice.go:56:10: invalid operation: carFuel += Liters(8) (mismatched types Gallons and Liters)
	// ./stractpractice.go:57:10: invalid operation: busFuel += Gallons(30) (mismatched types Liters and Gallons)
	carFuel += ToGallons(Liters(8.0))
	busFuel += ToLiters(Gallons(30.0))
	fmt.Printf("Car Fuel: %0.1f gallons\n", carFuel)
	fmt.Printf("Bus Fuel: %0.1f liters\n", busFuel)
}

func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (g Gallons) ToMilliliters() Milliliters {
	return Milliliters(g * 3785.41)
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}

func (m Milliliters) String() string {
	return fmt.Sprintf("%0.2f ml", m)
}
