package main

import (
	"flag"
	"fmt"
	"time"
)

var j = flag.Int("j", 0, "j")

func main() {
	flag.Parse()

	var i = 0
	for {
		fmt.Println("main println", i, *j)
		i++
		time.Sleep(1 * time.Second)
	}
}
