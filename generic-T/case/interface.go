package _case

import "fmt"

type ToString interface {
	String() string
}

func (u user) String() string {
	return fmt.Sprintf("user{Id: %d, Name: %s, Age: %d}", u.Id, u.Name, u.Age)
}

func (addr address) String() string {
	return fmt.Sprintf("address{Id : %d , ProvinceName : %s , CityName : %s}", addr.Id, addr.ProvinceName, addr.CityName)
}

type GetKey[T comparable] interface {
	any
	Get() T
}

func (u user) Get() int64 {
	return u.Id
}

func (addr address) Get() int {
	return addr.Id
}

func listToMap[k comparable, T GetKey[k]](list []T) map[k]T {
	m := make(map[k]T, 0)
	for _, v := range list {
		m[v.Get()] = v
	}
	return m
}

func InterfaceCase() {
	u := user{Id: 1, Name: "1", Age: 1}
	addr := address{Id: 1, ProvinceName: "bj", CityName: "bj"}
	fmt.Println(u.String())
	fmt.Println(addr.String())

	userList := []GetKey[int64]{user{Id: 1, Name: "1", Age: 1}, user{Id: 2, Name: "2", Age: 2}}
	addrList := []GetKey[int]{address{Id: 1, ProvinceName: "bj", CityName: "sy"}, address{Id: 2, ProvinceName: "zj", CityName: "hz"}}

	fmt.Println(listToMap(userList))
	fmt.Println(listToMap(addrList))

}
