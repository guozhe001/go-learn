package main

import "fmt"

type studect struct {
	name  string
	grade float64
}

func main() {
	var s studect
	s.name = "Alonzo Cole"
	s.grade = 92.3
	printInfo(s)
}

func printInfo(s studect) {
	fmt.Println("Name:", s.name)
	fmt.Printf("Grade: %0.1f\n", s.grade)
}
