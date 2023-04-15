package main

import (
	"fmt"
	"sync"
)

func main() {
	var pool sync.Pool
	pool.Put(1)
	pool.Put("hello1")
	pool.Put(3.14159)
	pool.Put("hello2")

	//i:=pool.Get()
	//fmt.Println(i)
	//i=pool.Get()
	//fmt.Println(i)

	for {
		val := pool.Get()
		if val == nil { //默认是nil
			break
		}
		//操作数据
		fmt.Println(val)
	}
}
