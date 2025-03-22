package _case

import "fmt"

type Dove struct{}

func (d *Dove) Eat() {
	fmt.Println("dove eat")
}
func (d *Dove) Run() {
	fmt.Println("dove run")
}
func (d *Dove) Speak() {
	fmt.Println("dove speak")
}

func NewDove() AnimalI {
	//return Dove{}  报错
	return &Dove{}
}
