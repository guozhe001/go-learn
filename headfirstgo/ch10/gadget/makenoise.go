package main

import "fmt"

type NoiseMaker interface {
	MakeSound()
}

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

func play(n NoiseMaker) {
	n.MakeSound()
	// 如果像下面这样直接把n类型断言为Robot，那么在传入Fan或其他类型的NoiseMaker时就会报错：
	// panic: interface conversion: main.NoiseMaker is main.Whistle, not main.Robot
	// var robot Robot = n.(Robot)
	// robot.Walk()
	robot, ok := n.(Robot)
	if ok {
		robot.Walk()
	}
}

func main() {
	play(Whistle("Toyco Canary"))
	play(Horn("Toyco Blaster"))
	play(Robot("Botco Ambler"))
}

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep Boop")
}

func (r Robot) Walk() {
	fmt.Println("Powering legs")
}
