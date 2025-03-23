package main

import (
	_case "godemo/function/case"
	"os"
	"os/signal"
)

func main() {
	_case.Closure()
	ch := make(chan os.Signal, 0)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
}
