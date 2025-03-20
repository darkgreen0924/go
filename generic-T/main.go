package main

import (
	"context"
	_case "godemo/generic-T/case"
	"os"
	"os/signal"
)

func main() {
	_case.Simple()
	_case.TTypeCase()
	_case.TTypeCase1()
	_case.InterfaceCase()

	// 不让主协程过早退出

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()
	<-ctx.Done()
}
