package main

import (
	"fmt"
	"sync"
)

var s3 string

var once sync.Once

func foo() {
	s3 = "hello world"
}
func main() {
	once.Do(foo)
	fmt.Println(s3) //一定会打印hello world
}
