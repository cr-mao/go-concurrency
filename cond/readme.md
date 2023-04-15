## sync.Cond


Go 标准库提供 Cond 原语的目的是，为等待 / 通知场景下的并发问题提供支持。Cond 通常应用于等待某个条件的一组 goroutine，等条件变为 true 的时候，其中一个 goroutine 或者所有的 goroutine 都会被唤醒执行。




对应的有3个常用方法，Wait，Signal，Broadcast。

```text 
1)	func (c *Cond) Wait()  wait  之前加上锁
该函数的作用可归纳为如下三点：
a)	阻塞等待条件变量满足	
b)	释放已掌握的互斥锁相当于cond.L.Unlock()。 注意：两步为一个原子操作。（原子操作，一起获得cpu ）
c)	当被唤醒时，Wait()函数返回时，解除阻塞并重新获取互斥锁。相当于cond.L.Lock()

2)	func (c *Cond) Signal()
	单发通知，给一个正等待（阻塞）在该条件变量上的goroutine（线程）发送通知。
3)	func (c *Cond) Broadcast()
广播通知，给正在等待（阻塞）在该条件变量上的所有goroutine（线程）发送通知。
```



条件变量的`Wait`方法主要做了四件事。

1. 把调用它的 goroutine（也就是当前的 goroutine）加入到当前条件变量的通知队列中。
2. 解锁当前的条件变量基于的那个互斥锁。
3. 让当前的 goroutine 处于等待状态，等到通知到来时再决定是否唤醒它。此时，这个 goroutine 就会阻塞在调用这个`Wait`方法的那行代码上。
4. 如果通知到来并且决定唤醒这个 goroutine，那么就在唤醒它之后重新锁定当前条件变量基于的互斥锁。自此之后，当前的 goroutine 就会继续执行后面的代码了。

