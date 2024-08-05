## waitgroup 


常见问题：

1. 计数器变负数
```go
var wg sync.WaitGroup 
wg.Add(10)
wg.Add(-10)
wg.Add(-1) //panic 
```

或者 Done()方法执行太多。
```go
var wg sync.WaitGroup 
wg.Add(1)
wg.Done()
wg.Done() //panic 
```


noCopy 辅助检查，vet工具 



用这里的并发原语，可以减少犯错
https://github.com/sourcegraph/conc 
