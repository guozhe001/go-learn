// 第356页冰箱贴练习
package main

import (
	"fmt"
	"log"
)

type Refrigerator []string

// Open 打开
func (r Refrigerator) Open() {
	fmt.Println("Opening refrigerator")
}

// Close 关闭
func (r Refrigerator) Close() {
	fmt.Println("Closing refrigerator")
}

// FindFood 找食物
func (r Refrigerator) FindFood(food string) error {
	r.Open()
	defer r.Close()
	if find(food, r) {
		fmt.Println("Found", food)
	} else {
		return fmt.Errorf("%s not found", food)
	}
	return nil
}

func find(item string, slice []string) bool {
	for _, sliceItem := range slice {
		if sliceItem == item {
			return true
		}
	}
	return false
}

func main() {
	fridge := Refrigerator{"Milk", "Pizza", "Salsa"}
	for _, food := range []string{"Milk", "Bananas"} {
		err := fridge.FindFood(food)
		if err != nil {
			log.Fatal(err)
		}
	}
}
