## mutex 



1. 检测是否有竞争
```shell
go run -race no_lock.go 
```

2. vet 静态分析 死锁问题
```shell 
go vet copy_err.go
```


3. 可重入锁

   github.com/petermattis/goid 获得goid

4. TryLock  

go在1.18中已经增加这个功能 

Mutex.TryLock：尝试锁定互斥锁，返回是否成功。

5. 安全队列


