package main

import "fmt"

type Appliance interface {
	TurnOn()
}

type Fan string

func (f Fan) TurnOn() {
	fmt.Println("Spinning")
}

// func (f Fan) TurnOn() error {
// 	fmt.Println("Spinning")
// 	return nil
// }
// 如果在Fan的TurnOn方法上增加返回参数，则Fan就就不是Appliance类型了，会报如下的错误：
// cannot use Fan("Windco Breeze") (constant "Windco Breeze" of type Fan) as Appliance value in assignment: wrong type for method TurnOn (have func() error, want func())compiler

type CoffeePot struct {
	name string
}

func (c CoffeePot) TurnOn() {
	fmt.Println("Heating Up")
}

func main() {
	var device Appliance
	device = Fan("Windco Breeze")
	device.TurnOn()
	device = CoffeePot{name: "LuxBrew"}
	device.TurnOn()
}
