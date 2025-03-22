package _case

import "fmt"

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
