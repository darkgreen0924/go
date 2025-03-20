package _case

import (
	"fmt"
)

type user struct {
	Id   int64
	Name string
	Age  int
}

type address struct {
	Id           int
	ProvinceName string
	CityName     string
}

func mapToList[k comparable, T any](mp Map[k, T]) []T {
	list := make([]T, len(mp))
	var i int
	for _, data := range mp {
		list[i] = data
		i++
	}
	return list
}

func myPrint[T any](ch chan T) {
	for data := range ch {
		fmt.Println(data)
	}
}

func TTypeCase() {
	userMp := make(Map[int, user], 0)

	userMp[1] = user{Id: 1, Name: "1", Age: 1}
	userMp[2] = user{Id: 2, Name: "2", Age: 1}
	userMp[3] = user{Id: 3, Name: "3", Age: 1}

	list := mapToList(userMp)

	ch := make(chan user)

	go func() {
		myPrint(ch)
	}()

	for _, data := range list {
		ch <- data
	}

}

func TTypeCase1() {
	userMp := make(Map[int, user], 0)

	userMp[1] = user{Id: 1, Name: "1", Age: 1}
	userMp[2] = user{Id: 2, Name: "2", Age: 1}
	userMp[3] = user{Id: 3, Name: "3", Age: 1}

	list := mapToList(userMp)

	ch := make(Chan[user])

	go func() {
		myPrint(ch)
	}()

	for _, data := range list {
		ch <- data
	}

}

type Chan[T any] chan T

type Map[k comparable, T any] map[k]T

type List[T any] []T
