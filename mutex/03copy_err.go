package main

import (
	"fmt"
	"sync"
)

// 锁 copy 带来的死锁

type Counter struct {
	sync.Mutex
	count int
}

func foo(c Counter) {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo")
}

func main() {
	var c Counter
	c.Lock()
	defer c.Unlock()
	c.count++
	foo(c)
}
