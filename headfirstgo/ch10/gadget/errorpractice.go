package main

import "fmt"

type ComedyError string

func (c ComedyError) Error() string {
	return string(c)
}

func main() {
	var err error
	err = ComedyError("error!")
	fmt.Println(err)

	err = checkTemperature(121.2, 100)
	fmt.Println(err)
}

type OverHeatError float64

func (o OverHeatError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degrees!", o)
}

func checkTemperature(actual float64, safe float64) error {
	if actual-safe > 0 {
		return OverHeatError(actual)
	}
	return nil
}
