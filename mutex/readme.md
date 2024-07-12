## mutex

1. 检测是否有竞争

```shell
go run -race 01no_lock.go 
```

2. 加索访问

3. vet 静态分析 死锁问题

```shell 
go vet 03copy_err.go
```

4. 可重入锁
   goid 实现
5. 可重入锁
   token 方式实现
6. TryLock

    go在1.18中已经增加这个功能

    Mutex.TryLock：尝试锁定互斥锁，返回是否成功。

7. 安全队列

8. double check 常见写法

```shell
go run 008double_check.go
```

9. 并发 map，基于mutex 实现，后面有读写锁

10. 泛型 mutex 实现

mutex使用的注意事项:

- 使用零值，不需要显示初始化 mutex。
- 尽量少写 if ，else 中去释放锁，能用 defer 尽量用 defer。因为 go1.14做了内敛方式优化，取代之前的生成 defer 对象到
  defer 链中。

// 新版本 go，在代码上面增加此行，表示race 跳过检测
// go:build !race
// +build !race //老版本，如果要兼容，加 2 条






