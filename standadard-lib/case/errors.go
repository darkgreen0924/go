package _case

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type CusError struct {
	Code string
	Msg  string
	Time time.Time
}

func (e CusError) Error() string {
	return fmt.Sprintf("%s(%s,time=%s)", e.Code, e.Msg, e.Time.Format("2006-01-02 15:04:05"))
}

func getCusError(code, msg string) error {
	return CusError{
		Code: code,
		Msg:  msg,
		Time: time.Now(),
	}
}
func ErrorsCase() {
	err := errors.New("程序发生了错误")
	fmt.Println(err)

	var a, b = -1, -2
	res, err := sum(a, b)
	fmt.Println(res, err)
	if err != nil {
		log.Println(err)
		cusErr, ok := err.(CusError)
		if ok {
			fmt.Println("打印自定义错误信息", cusErr.Code, cusErr.Msg, cusErr.Time)
		}
	}

}
func sum(a, b int) (int, error) {
	if a < 0 && b < 0 {
		return 0, getCusError("500", "两数求和不能同时小于0")
	}
	return a + b, nil
}
