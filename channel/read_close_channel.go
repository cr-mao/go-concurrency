package main

import "fmt"

var ch1 = make(chan struct{}, 1)

var s1 string

func f2() {
	s1 = "hello world"
	close(ch1)
}

func main() {
	go f2()
	fmt.Println(<-ch1) // 读到 0值: {}
	fmt.Println(<-ch1) // 读到 0值: {}
	print(s1)          //一定打印hello，world， channel第n个写 第一个happens before n个读的完成
}
