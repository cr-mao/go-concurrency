## Go Concurrency

go 并发相关总结

- sync.Mutex
- sync.RWMutex
- sync.Once
- sync.Cond
- sync.Map
- sync.WaitGroup
- sync.Pool
- channel
- context
- 内存模型


如何选择
1. 共享资源的并发访问使用传统并发原语；
2. 复杂的任务编排和消息传递使用 Channel；
3. 消息通知机制使用 Channel，除非只想 signal 一个 goroutine，才使用 Cond；
4. 简单等待所有任务的完成用 WaitGroup，也有 Channel 的推崇者用 Channel，都可以；
5. 需要和 Select 语句结合，使用 Channel；
6. 需要和超时配合时，使用 Channel 和 Context。

[go中happens-before 保证](memory/readme.md)



## links

- 《极客时间-go并发编程实战课》
