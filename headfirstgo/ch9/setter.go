package main

import (
	"fmt"
	"log"

	"github.com/guozhe001/go-learn/headfirstgo/ch9/calendar"
)

func main() {
	event := calendar.Event{}
	if err := event.SetYear(2021); err != nil {
		log.Fatal(err)
	}
	if err := event.SetMonth(1); err != nil {
		log.Fatal(err)
	}
	if err := event.SetDay(32); err != nil {
		log.Fatal(err)
	}
	if err := event.SetTitle("learn"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())
	fmt.Println(event.Title())
}
