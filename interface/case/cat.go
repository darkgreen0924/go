package _case

import "fmt"

type Cat struct {
	Animal
}

func NewCat() AnimalI {
	return &Cat{}
}

func (cat Cat) Eat() {
	fmt.Println("猫吃老鼠")
}
