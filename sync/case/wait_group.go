package _case

import (
	"fmt"
	"sync"
	"time"
)

func WaitGroupCase() {
	start := time.Now()
	var a, b = 100, 100
	for i := 0; i < 8000000000; i++ {
		multi(a, b)
	}
	t := time.Since(start)
	fmt.Println(t)

	wg := sync.WaitGroup{}
	start = time.Now()

	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000000000; j++ {
				multi(a, b)
			}
		}()
	}
	t = time.Since(start)
	wg.Wait()
	fmt.Println(t)

}

func WaitGroupCase2() {
	ch := make(chan []int, 1000)
	wg1 := &sync.WaitGroup{}
	wg2 := &sync.WaitGroup{}
	start := time.Now()
	wg1.Add(1)
	go func() {
		defer wg1.Done()
		var i int = 0
		for item := range ch {
			fmt.Println(multi(item[0], item[1]))
			i++
		}
		time.Sleep(3 * time.Second)
		fmt.Println("数据条数", i)
	}()

	for i := 1; i <= 2; i++ {
		wg2.Add(1)
		wg1.Add(1)

		go func(wg *sync.WaitGroup) {
			defer wg1.Done()
			defer wg.Done()
			for j := 0; j < 500; j++ {
				ch <- []int{i, j}
			}
		}(wg2)
	}

	wg2.Wait()
	close(ch)
	wg1.Wait()

	time.Since(start)
}
func multi(a, b int) int {
	return a * b
}
