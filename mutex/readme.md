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

