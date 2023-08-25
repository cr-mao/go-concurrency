## go内存模型:Go如何保证并发读写的顺序


由于指令重排，代码并不一定会按照你写的顺序执行.

```go
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
```

## happens before

**在一个 goroutine 内部，程序的执行顺序和它们的代码指定的顺序是一样的，即使编译器或者 CPU 重排了读写顺序，从行为上来看，也和代码指定的顺序一样**

但是，对于另一个 goroutine 来说，重排却会产生非常大的影响。 Go 只保证 goroutine 内部重排对读写的顺序没有影响.


Go 内存模型通过 happens-before 定义两个事件（读、写 action）的顺序：如果事件 e1  happens before 事件 e2，那么，我们就可以说事件 e2 在事件 e1 之后发生（happens after）。如果 e1 不是 happens before e2， 同时也不 happens after e2，那么，我们就可以说事件 e1 和 e2 是同时发生的。

如果要保证对变量v的读操作r 能够观察到一个对变量v的写操作w,并且r只能观察到w对变量的写，没有其他对变量v的写。要满足2个条件
- 1. w happens before r；
- 其它对 v 的写操作（w2、w3、w4, …） 要么 happens before w，要么 happens after r，绝对不会和 w、r 同时发生，或者是在它们之间发生。


在单个的 goroutine 内部， happens-before 的关系和代码编写的顺序是一致的。

下面一定会依次打印1，2，3 
```go
func foo(){
	var a = 1
	var b = 2
	var c = 3
	println(a)
	println(b)
	println(c)
}
```

## Go 语言中保证的 happens-before 关系


除了单个 goroutine 内部提供的 happens-before 保证，Go 语言中还提供了一些其它的 happens-before 关系的保证。


### init函数

包级别的变量在同一个文件中是按照声明顺序逐个初始化的，除非初始化它的时候依赖其它的变量。同一个包下的多个文件，会按照文件名的排列顺序进行初始化。

具体怎么对这些变量进行初始化呢？Go 采用的是依赖分析技术。不过，依赖分析技术保证的顺序只是针对同一包下的变量，而且，只有引用关系是本包变量、函数和非接口的方法，才能保证它们的顺序性。

```go
var (
	a = 9  
	b = f()   //4
	c = f()   //5
	d = 3    //全部初始化后 =5
)

func f() int {
	d++
	return d
}
```

同一个包下可以有多个 init 函数，但是每个文件最多只能有一个 init 函数，多个 init 函数按照它们的文件名顺序逐个初始化。

包的倒入顺序


### goroutine 

启动 goroutine 的 go 语句的执行，一定 happens before 此 goroutine 内的代码执行。

在下面的代码中，第 8 行 a 的赋值和第 9 行的 go 语句是在同一个 goroutine 中执行的，所以，在主 goroutine 看来，第 8 行肯定 happens before 第 9 行，又由于刚才的保证，第 9 行子 goroutine 的启动 happens before 第 4 行的变量输出，那么，我们就可以推断出，第 8 行 happens before 第 4 行。也就是说，在第 4 行打印 a 的值的时候，肯定会打印出“hello world”

```go
var a string 

func f(){
	print(a)       //第4行
}

func hello(){
	a = "hello world"  //第8行
	go f()            // 第9行
}
```


### channel

channel中保证happens before 

1.往 Channel 中的发送操作，happens before 从该 Channel 接收相应数据的动作完成之前，即第 n 个 send 一定 happens before 第 n 个 receive 的完成。

2. close 一个 Channel 的调用，肯定 happens before 从关闭的 Channel 中读取出一个零值。

3. 对于 unbuffered 的 Channel，也就是容量是 0 的 Channel，从此 Channel 中读取数据的调用一定 happens before 往此 Channel 发送数据的调用完成。

4. 如果 Channel 的容量是 m（m>0），那么，第 n 个 receive 一定 happens before 第 n+m 个 send 的完成。


### Mutex、RWMutex

对于互斥锁 Mutex m 或者读写锁 RWMutex m，有 3 条 happens-before 关系的保证。
1. 第 n 次的 m.Unlock 一定 happens before 第 n+1 m.Lock 方法的返回；
2. 对于读写锁 RWMutex m，如果它的第 n 个 m.Lock 方法的调用已返回，那么它的第 n 个 m.Unlock 的方法调用一定 happens before 任何一个 m.RLock 方法调用的返回，只要这些 m.RLock 方法调用 happens after 第 n 次 m.Lock 的调用的返回。这就可以保证，只有释放了持有的写锁，那些等待的读请求才能请求到读锁。
3. 对于读写锁 RWMutex m，如果它的第 n 个 m.RLock 方法的调用已返回，那么它的第 k （k<=n）个成功的 m.RUnlock 方法的返回一定 happens before 任意的 m.RUnlockLock 方法调用，只要这些 m.Lock 方法调用 happens after 第 n 次 m.RLock。

对于读写锁l的l.RLock的调用，如果存在一个n，这次的 l.RLock 调用 happens after 第 n 次的 l.Unlock，那么，和这个 RLock 相对应的 l.RUnlock 一定 happens before 第 n+1 次 l.Lock。意思是，读写锁的 Lock 必须等待既有的读锁释放后才能获取到。


### WaitGroup 

对于一个 WaitGroup 实例 wg，在某个时刻 t0 时，它的计数值已经不是零了，假如 t0 时刻之后调用了一系列的 wg.Add(n) 或者 wg.Done()，并且只有最后一次调用 wg 的计数值变为了 0，那么，可以保证这些 wg.Add 或者 wg.Done() 一定 happens before t0 时刻之后调用的 wg.Wait 方法的返回。

Wait 方法等到计数值归零之后才返回


### Once

对于 once.Do(f) 调用，f 函数的那个单次调用一定 happens before 任何 once.Do(f) 调用的返回


### atomic

对于 Go 1.15 的官方实现来说，可以保证使用 atomic 的 Load/Store 的变量之间的顺序性。

```go
package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

func main() {
	var a, b int32 = 0, 0

	go func() {
		atomic.StoreInt32(&a, 1)
		atomic.StoreInt32(&b, 1)
	}()

	for atomic.LoadInt32(&b) == 0 {
		runtime.Gosched()
	}
	fmt.Println(atomic.LoadInt32(&a)) // b=1了， a 一定也是1 。
}
```
