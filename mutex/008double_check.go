/**
User: cr-mao
Date: 2023/8/5 06:29
Email: crmao@qq.com
Desc: 008double_check.go
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

type User struct {
	sync.RWMutex
	username string
}

var user = &User{}

func main() {
	wait := &sync.WaitGroup{}
	wait.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			loadOrStore("cr-mao", wait)

		}()
	}
	wait.Wait()

	time.Sleep(1000 * time.Millisecond)
}

func loadOrStore(username string, wg *sync.WaitGroup) string {
	defer wg.Done()
	defer func() {
		fmt.Println(user.username)
	}()
	if user.username != "" {
		return user.username
	}
	user.RLock()
	if user.username != "" {
		return user.username
	}
	user.RUnlock()

	user.Lock()
	user.username = username
	user.Unlock()
	return user.username
}
