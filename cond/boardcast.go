package main

import (
	"log"
	"sync"
	"time"
)

/*
条件变量的`Signal`方法和`Broadcast`方法有哪些异同？

条件变量的`Signal`方法和`Broadcast`方法都是被用来发送通知的，不同的是，前者的通知只会唤醒一个因此而等待的 goroutine，而后者的通知却会唤醒所有为此等待的 goroutine。

条件变量的`Wait`方法总会把当前的 goroutine 添加到通知队列的队尾，而它的`Signal`方法总会从通知队列的队首开始，查找可被唤醒的 goroutine。所以，因`Signal`方法的通知，而被唤醒的 goroutine 一般都是最早等待的那一个。

这两个方法的行为决定了它们的适用场景。如果你确定只有一个 goroutine 在等待通知，或者只需唤醒任意一个 goroutine 就可以满足要求，那么使用条件变量的`Signal`方法就好了。

否则，使用`Broadcast`方法总没错，只要你设置好各个 goroutine 所期望的共享资源状态就可以了。

此外，再次强调一下，与`Wait`方法不同，条件变量的`Signal`方法和`Broadcast`方法并不需要在互斥锁的保护下执行。恰恰相反，我们最好在解锁条件变量基于的那个互斥锁之后，再去调用它的这两个方法。这更有利于程序的运行效率。

最后，请注意，条件变量的通知具有即时性。也就是说，如果发送通知的时候没有 goroutine 为此等待，那么该通知就会被直接丢弃。在这之后才开始等待的 goroutine 只可能被后面的通知唤醒。
*/
func main() {
	// mailbox 代表信箱。
	// 0代表信箱是空的，1代表信箱是满的。
	var mailbox uint8
	// lock 代表信箱上的锁。
	var lock sync.Mutex
	// sendCond 代表专用于发信的条件变量。
	sendCond := sync.NewCond(&lock)
	// recvCond 代表专用于收信的条件变量。
	recvCond := sync.NewCond(&lock)

	// send 代表用于发信的函数。
	send := func(id, index int) {
		lock.Lock()
		for mailbox == 1 {
			sendCond.Wait()
		}
		log.Printf("sender [%d-%d]: the mailbox is empty.",
			id, index)
		mailbox = 1
		log.Printf("sender [%d-%d]: the letter has been sent.",
			id, index)
		lock.Unlock()
		recvCond.Broadcast()
	}

	// recv 代表用于收信的函数。
	recv := func(id, index int) {
		lock.Lock()
		for mailbox == 0 {
			recvCond.Wait()
		}
		log.Printf("receiver [%d-%d]: the mailbox is full.",
			id, index)
		mailbox = 0
		log.Printf("receiver [%d-%d]: the letter has been received.",
			id, index)
		lock.Unlock()
		sendCond.Signal() // 确定只会有一个发信的goroutine。
	}

	// sign 用于传递演示完成的信号。
	sign := make(chan struct{}, 3)
	max := 6
	go func(id, max int) { // 用于发信。
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= max; i++ {
			time.Sleep(time.Millisecond * 500)
			send(id, i)
		}
	}(0, max)
	go func(id, max int) { // 用于收信。
		defer func() {
			sign <- struct{}{}
		}()
		for j := 1; j <= max; j++ {
			time.Sleep(time.Millisecond * 200)
			recv(id, j)
		}
	}(1, max/2)
	go func(id, max int) { // 用于收信。
		defer func() {
			sign <- struct{}{}
		}()
		for k := 1; k <= max; k++ {
			time.Sleep(time.Millisecond * 200)
			recv(id, k)
		}
	}(2, max/2)

	<-sign
	<-sign
	<-sign
}
