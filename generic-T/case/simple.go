package _case

import "fmt"

func Simple() {
	var a, b = 1, 2
	var c, d = 1.1, 2.1
	maxInt := getMaxValueGenericT(a, b)
	maxFloat := getMaxValueGenericT(c, d)

	fmt.Println("maxInt=", maxInt)
	fmt.Println("maxFloat=", maxFloat)

	var e, f MyInt64 = 1, 2
	maxCusTomT := getMaxValueCusNumT(e, f)
	fmt.Println("maxCusTomT=", maxCusTomT)

	var h, i int = 1, 2
	fmt.Println("getBuiltInComparable=", getBuiltInComparable(h, i))
}
func getMaxValueInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMaxValueFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// 指定类型
func getMaxValueGenericT[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

//func func1[T any](case, b T) T {

//}

// 若是指针的话则需要这么写
//func func2[T interface{ *int | *float64 }](case, b T) T {
//
//}

type CusNumT interface {
	// 多行取交集
	// |并集
	// ~ 衍生类型
	uint8 | int32 | float64 | ~int64
}

// 衍生类型 会用到类型转换
type MyInt64 int64

// 别名 不会进行类型转换
type MyInt32 = int32

func getMaxValueCusNumT[T CusNumT](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// 内置类型 any comparable
func getBuiltInComparable[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}
