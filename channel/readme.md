## channel


```text
  channel <- value      // 
    <-channel             //接收并将其丢弃
    x := <-channel        //从channel中接收数据，并赋值给x
    x, ok := <-channel    //功能同上，同时检查通道是否已关闭或者是否为空 , 同时关闭和为空 才返回false
```

```go
//只是清空 chan 
for range ch1 {
	
}
```


空(nil) 读写阻塞， 写关闭异常，读关闭空零

panic场景：
- close 已经close 的chan 
- send 已经close的chan
- close ni的chan 


# todo 各种模式



