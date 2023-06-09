package main

import (
	"fmt"
	"sync"
)

func main() {
	var count = 0
	// 使用WaitGroup 等待10个goroutine 完成
	var wg sync.WaitGroup
	var lock sync.Mutex

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				lock.Lock()
				count++
				lock.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("count值是:", count)
}
