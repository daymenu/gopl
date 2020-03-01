# 基于共享变量的并发

## 竞态条件

只要有两个或两个以上的goroutine同时访问一个变量，并且其中至少一个是写操作就会发生数据竞争。

避免竞争的方法：

1. 不要去写变量
2. 避免从多个goroutine访问变量
3. 加锁，同一时刻只有各一个在访问

```go
mu := sync.rwmutex{}
mu.rlock()
mu.runlock()
mu1 := sync.mutex{}
mu1.lock()
mu1.unlock()
```

## 内存同步

每个cpu都会有自己的缓存，只有channel通信或者互斥量操作，才会将cache提交到主存，其他的goroutine才能读取到变化。

## 一次性初始化锁

```go
once := sync.Once{}
once.Do(func() {
    fmt.Println(1)
})
```

## 数据竞争条件检测

go build 、go test、 go run 后面加 -race 的flag

end
