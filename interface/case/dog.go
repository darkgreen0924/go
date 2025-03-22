package _case

import "fmt"

type Dog struct {
	Animal
}

func (dog Dog) Eat() {
	fmt.Println("dog eat")
}

func (dog Dog) Speak() {
	fmt.Println("dog speak")
}

func (dog Dog) Run() {
	fmt.Println("dog run")
}

func NewDog() AnimalI {
	return &Dog{}
	//return Dog{}
}
