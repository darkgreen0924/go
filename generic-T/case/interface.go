package _case

import "fmt"

type ToString interface {
	String() string
}

func (u user) String() string {
	return fmt.Sprintf("user{Id: %s, Name: %s, Age: %d}", u.Id, u.Name, u.Age)
}

func (addr address) String() string {
	return fmt.Sprintf("address{Id : %s , ProvinceName : %s , CityName : %s}", addr.Id, addr.ProvinceName, addr.CityName)
}

func InterfaceCase() {
	u := user{Id: "1", Name: "1", Age: 1}
	fmt.Println(u.String())
}
