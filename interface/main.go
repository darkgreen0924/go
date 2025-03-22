package main

import _case "godemo/interface/case"

func main() {
	cat := _case.NewCat()
	animalLife(cat)
	dog := _case.NewDog()
	animalLife(dog)
	dove := _case.NewDove()
	animalLife(dove)
}

func animalLife(a _case.AnimalI) {
	a.Speak()
	a.Eat()
	a.Run()
}
