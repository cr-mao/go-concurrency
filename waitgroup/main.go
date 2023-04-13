package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func(w *sync.WaitGroup) {
		defer w.Done()
		//业务
		fmt.Println("do business 1 ")
		time.Sleep(time.Second)
	}(&wg)

	go func(w *sync.WaitGroup) {
		defer w.Done()
		//业务
		fmt.Println("do business 2 ")
		time.Sleep(time.Second)
	}(&wg)
	wg.Wait()
}
