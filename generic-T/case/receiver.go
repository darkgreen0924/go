package _case

import "fmt"

type MyStruct[T interface{ *int | *string }] struct {
	Name string
	Data T
}

// 泛型 receiver
func (ms MyStruct[T]) Get() T {
	return ms.Data
}

func ReceiverCase() {
	data := "abc"
	var a = MyStruct[*string]{Name: "a", Data: &data}
	fmt.Println(&a)
}
