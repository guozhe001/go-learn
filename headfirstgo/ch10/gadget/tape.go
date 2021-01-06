package main

import (
	"fmt"
)

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(name string) {
	fmt.Println("Playing:", name)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped")
}

type TapeRecorder struct {
	Microphones int
}

// 如果像下面这样定义一个struct，会报如下的错误
// ./tape.go:62:8: type TapeRecorder1 is not an expression
// type TapeRecorder1 struct {

// }

// func (t TapeRecorder1) Play(name string) {
// 	fmt.Println("Playing:", name)
// }

// func (t TapeRecorder1) Record() {
// 	fmt.Println("Recording")
// }

// func (t TapeRecorder1) Stop() {
// 	fmt.Println("Stopped")
// }

func (t TapeRecorder) Play(name string) {
	fmt.Println("Playing:", name)
}

func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopped")
}

func playList(divice Player, songs []string) {
	for _, song := range songs {
		divice.Play(song)
	}
	divice.Stop()
}

func main() {
	var payer Player = TapePlayer{}
	mixtape := []string{"a", "b", "c"}
	playList(payer, mixtape)
	TryOut(TapeRecorder{})
	TryOut(TapePlayer{})
}

type Player interface {
	Play(song string)
	Stop()
}

// TryOut 练习类型断言
func TryOut(player Player) {
	player.Play("try")
	player.Stop()
	// 如果直接使用类型转换的方式把接口转换成某个类型，会报如下的错误
	// ./tape.go:76:14: cannot convert player (type Player) to type TapeRecorder: need type assertion
	// TapeRecorder(player).Record()
	// 使用类型断言的方式转换
	tapeRecorder, ok := player.(TapeRecorder)
	if ok {
		tapeRecorder.Record()
	}
}
