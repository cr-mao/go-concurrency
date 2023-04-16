package main

//重排 ，可见性

var a, b int

func f() {
	a = 1 //w之前的写操作
	b = 2 // 写操作 w
}
func g() {
	print(b) //读操作r
	print(a) //哪怕 上面打印2 ,这里a 也可能是0
}
func main() {
	go f() //g1
	g()    //g2
}
