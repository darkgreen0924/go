package _case

import (
	"fmt"
	"time"
)

// 协程间通信
func Communication() {
	// 定义一个可读可写通道
	ch := make(chan int, 0)
	go communicationF1(ch)
	go communicationF2(ch)
}

// F1 定义一个接受只写通道
func communicationF1(ch chan<- int) {
	for i := range 100 {
		ch <- i
	}
}

// F2定义一个只读通道
func communicationF2(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

// 并发情况下
func ConcurrentSyncCase() {
	// 带缓冲的管道
	ch := make(chan int, 10)
	go func() {
		for i := range 100 {
			ch <- i
		}
	}()

	go func() {
		for i := range 100 {
			ch <- i
		}
	}()

	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()
}

func NoticeAndMultiplexingCase() {
	ch := make(chan int, 0)
	strCh := make(chan string, 0)
	done := make(chan struct{}, 0)
	go noticeAndMultiplexingF1(ch)
	go noticeAndMultiplexingF2(strCh)
	go noticeAndMultiplexingF3(ch, strCh, done)

	time.Sleep(1 * time.Second)
	close(done)
}

func noticeAndMultiplexingF1(ch chan<- int) {
	for i := range 100 {
		ch <- i
	}
}

func noticeAndMultiplexingF2(strCh chan<- string) {
	for i := range 100 {
		strCh <- fmt.Sprintf("str %d", i)
	}
}

func noticeAndMultiplexingF3(ch <-chan int, strCh <-chan string, done <-chan struct{}) {
	i := 0
	for {
		select {
		case i = <-ch:
			fmt.Println(i)
		case str := <-strCh:
			fmt.Println(str)
		case <-done:
			fmt.Println("done")
			return
		default:
			fmt.Println("default")
		}
		i++
		fmt.Println("累计执行次数", i)
	}
}
