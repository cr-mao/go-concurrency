package main

import (
	"fmt"
	"runtime"
	"sync"
)

// sync.Pool 对象缓存

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("create a new object")
			return 100
		},
	}

	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)  //私有对象
	runtime.GC() //gc 会请空 sync.pool 缓存的对象
	v1, _ := pool.Get().(int)
	fmt.Println(v1) // 100 or 3
	v2, _ := pool.Get().(int)
	fmt.Println(v2) // 100
}

/*
➜  pool git:(main) ✗ go run gc_pool.go
create a new object
100
3
create a new object
100

➜  pool git:(main) ✗ go run gc_pool.go
create a new object
100
create a new object
100
create a new object
100


*/
