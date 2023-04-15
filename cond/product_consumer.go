package main

import "fmt"
import "sync"
import "math/rand"
import "time"

var cond sync.Cond // 创建全局条件变量

// 生产者
func producer(out chan<- int, idx int) {
	for {
		cond.L.Lock()       // 条件变量对应互斥锁加锁
		for len(out) == 3 { // 产品区满 等待消费者消费 ，必须用for ，if的话回不去了，代码走下去，写入channel 阻塞，通知不到消费者了
			cond.Wait() // 挂起当前协程， 等待条件变量满足，被消费者唤醒
			//   阻塞,解锁，    当被唤醒时加锁
		}
		num := rand.Intn(1000) // 产生一个随机数
		out <- num             // 写入到 channel 中（生产）
		fmt.Printf("%dth 生产者，产生数据 %3d, 公共区剩余%d个数据\n", idx, num, len(out))
		cond.L.Unlock()         // 生产结束，解锁互斥锁
		cond.Signal()           // 唤醒 阻塞的 消费者
		time.Sleep(time.Second) // 生产完休息一会，给其他协程执行机会
	}
}

//消费者
func consumer(in <-chan int, idx int) {
	for {
		cond.L.Lock()      // 条件变量对应互斥锁加锁（与生产者是同一个）
		for len(in) == 0 { // 产品区为空 等待生产者生产
			cond.Wait() // 挂起当前协程， 等待条件变量满足，被生产者唤醒
		}
		num := <-in // 将 channel 中的数据读走 （消费）
		fmt.Printf("---- %dth 消费者, 消费数据 %3d,公共区剩余%d个数据\n", idx, num, len(in))
		cond.L.Unlock()                    // 消费结束，解锁互斥锁
		cond.Signal()                      // 唤醒阻塞的生产者
		time.Sleep(time.Millisecond * 500) //消费完 休息一会，给其他协程执行机会
	}
}
func main() {
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子
	quit := make(chan bool)          // 创建用于结束通信的 channel

	product := make(chan int, 3) // 产品区（公共区）使用channel 模拟
	cond.L = new(sync.Mutex)     // 创建互斥锁和条件变量

	for i := 0; i < 5; i++ { // 5个消费者
		go producer(product, i+1)
	}
	for i := 0; i < 3; i++ { // 3个生产者
		go consumer(product, i+1)
	}
	<-quit // 主协程阻塞 不结束
}
