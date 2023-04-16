package main

var ch = make(chan struct{}, 1)

var s string

func f1() {
	s = "hello world"
	ch <- struct{}{}
}

func main() {
	go f1()
	<-ch
	print(s) //一定打印hello，world， channel第n个写 第一个happens before n个读的完成
}
