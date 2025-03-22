package _case

import (
	"errors"
	"fmt"
	"reflect"
)

func ReflectCase() {
	type user struct {
		Name string
		Age  int
	}
	var u user = user{Name: "John", Age: 25}
	var u1 user
	err := copy(&u1, &u)
	fmt.Println(err, u1)
}
func copy(dest any, src any) error {
	srcType := reflect.TypeOf(src)
	srcValue := reflect.ValueOf(src)
	fmt.Println(srcType.Elem(), srcValue.Elem())

	// 如果该类型是指针类型 则获取值
	if srcValue.Kind() == reflect.Ptr {
		srcType = srcType.Elem()
		srcValue = srcValue.Elem()
	}

	destType := reflect.TypeOf(dest)
	destValue := reflect.ValueOf(dest)
	if destType.Kind() != reflect.Ptr {
		return errors.New("dest is not a pointer")
	}
	destType = destType.Elem()
	destValue = destValue.Elem()
	if srcValue.Kind() != reflect.Struct {
		return errors.New("src is not a struct or struct pointer")
	}
	if destType.Kind() != reflect.Struct {
		return errors.New("dest is not a struct or struct pointer")
	}

	destObject := reflect.New(destType)
	for i := 0; i < destType.NumField(); i++ {
		destField := destType.Field(i)
		if srcField, ok := srcType.FieldByName(destField.Name); ok {
			if srcField.Type != destField.Type {
				continue
			}
			srcValue := srcValue.FieldByName(srcField.Name)
			destObject.Elem().FieldByName(destField.Name).Set(srcValue)
		}
	}
	destValue.Set(destObject.Elem())
	return nil
}
