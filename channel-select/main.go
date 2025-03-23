package main

import (
	_case "godemo/channel-select/case"
	"os"
	"os/signal"
)

func main() {
	//_case.Communication()
	//_case.ConcurrentSyncCase()
	_case.NoticeAndMultiplexingCase()
	ch := make(chan os.Signal, 0)
	signal.Notify(ch, os.Interrupt, os.Kill)
	//阻塞
	<-ch
}
