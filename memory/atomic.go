package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

func main() {
	var a, b int32 = 0, 0

	go func() {
		atomic.StoreInt32(&a, 1)
		atomic.StoreInt32(&b, 1)
	}()

	for atomic.LoadInt32(&b) == 0 {
		runtime.Gosched()
	}
	fmt.Println(atomic.LoadInt32(&a)) // b=1了， a 一定也是1 。
}
