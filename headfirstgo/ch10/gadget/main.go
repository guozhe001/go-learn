package main

import "github.com/guozhe001/go-learn/headfirstgo/ch10/gadget"

func playList(divice gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		divice.Play(song)
	}
	divice.Stop()
}

func main() {
	payer := gadget.TapePlayer{}
	mixtape := []string{"a", "b", "c"}
	playList(payer, mixtape)
}
