package _case

import (
	"fmt"
	"os"
)

func FmtCase() {
	// 打印到标准输出
	fmt.Println("今天是2025年3月21号22:08")
	// 格式化，并打印到标准输出
	fmt.Printf("%s天气很好\n", "今天")
	// 格式化
	str := fmt.Sprintf("%s天气很好\n", "today")
	// 输出到io流
	fmt.Fprint(os.Stdout, str)
}
