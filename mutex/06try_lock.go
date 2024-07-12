package main

import (
	"fmt"
	"sync"
)

type Counter2 struct {
	sync.Mutex
	count int
}

func foo2(c Counter2) {
	if c.TryLock() {
		c.count++
		c.Unlock()
		fmt.Println("try lock  return true ")
	} else {
		fmt.Println("try lock return false")
	}
}

func main() {
	var c Counter2
	c.Lock()
	c.count++
	foo2(c) // fmt.Println("try lock return false")
	c.Unlock()
	fmt.Println(c.count) // 1
}
