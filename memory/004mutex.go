/**
User: cr-mao
Date: 2023/8/25 11:01
Email: crmao@qq.com
Desc: 004mutex.go
*/
package main

import "sync"

var mutx = sync.Mutex{}

var s4 string

func foo4() {
	s4 = "hello world"
	mutx.Unlock()
}

func main() {
	mutx.Lock()
	go foo4()
	mutx.Lock()
	print(s4) //一定会打印错hello world
}
