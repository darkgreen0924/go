package _case

import (
	"fmt"
)

type AnimalI interface {
	Speak()

	Eat()

	Run()
}

type Animal struct{}

func (a Animal) Speak() {
	fmt.Println("Animal each func 默认实现 Speak")
}

func (a Animal) Eat() {
	fmt.Println("Animal each func 默认实现 Eat")
}

func (a Animal) Run() {
	fmt.Println("Animal each func 默认实现 Run")
}

func init() {
	a := Animal{}
	a.Speak()
	a.Eat()
	a.Run()
}
