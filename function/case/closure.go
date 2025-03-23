package _case

import "fmt"

func Closure() {
	t := tool()

	t()
	t()
}
func tool() func() int {

	var a = 1
	var b = 2
	var c = 3
	return func() int {
		fmt.Println("a=", a, "b=", b, "c=", c)
		a = b
		b = c
		return a + b
	}
}
