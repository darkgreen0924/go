package _case

import "fmt"

// defer case
func DeferCase() {
	defer func() {
		fmt.Println("defer defer")
	}()
	fmt.Println("函数执行")
}

// defer初始化值
func DeferCase2() {
	var i = 1

	defer func(j int) {
		fmt.Println("j:", j)
	}(i + 1)

	defer func() {
		fmt.Println("i:", i+1)
	}()
	i = i + 99
	fmt.Println("i:", i)

}

func ExceptionCase() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("处理异常", err)
		}
	}()

	fmt.Println("开始执行方法")
	panic("抛出异常exception")
	fmt.Println("执行结束")
}
