package _case

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func EncodingCase() {
	type user struct {
		ID   int64
		Name string
		Age  int
	}
	u := user{ID: 1, Name: "John Doe", Age: 25}

	// 序列化
	bytes, err := json.Marshal(u)
	fmt.Println(bytes, err)

	u1 := user{}
	err = json.Unmarshal(bytes, &u1)
	fmt.Println(u1, err)

	// base64编解码
	str := base64.StdEncoding.EncodeToString(bytes)
	fmt.Println(str)

	bytes1, err := base64.StdEncoding.DecodeString(str)
	fmt.Println(bytes1, err)

	// 16进制编解码
	str1 := hex.EncodeToString(bytes)
	fmt.Printf("str1: %v\n", str1)

	bytes2, err := hex.DecodeString(str1)
	fmt.Println(bytes2, err)
}
