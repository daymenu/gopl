# goroutines和channels

## goroutines

> 在go语言中每一个并发执行的单元叫作一个goroutine。当一个主函数启动时，其主函数即在一个单独的goroutine中运行，我们叫它main goroutine。新的goroutine会用go语句来创建。

```go

func f() {
    fmt.Println("hello")
}

// 启动一个goroutine
go f()
```

## channels

> 如果说goroutine是go语言程序的并发体的话，那么channels就是它们之间的通信机制。一个channels是一个通信机制，它可以通过一个goroutine向另一个goroutine发送值信息。每个channel都友一个特殊的类型，也就是可以发送的数据类型。一个可以发送int类型的channel一般写为chan int。

```go
// make(chan T, cap)
ch := make(chan int)

单向channel
<- chan int //只接受
chan <-int // 只发送

func counter(naturals chan<- int) {
    for i := 0; i < 10; i++ {
        naturals <- i
    }
    close(naturals)
}

func squarer(squares chan<- int, naturals <-chan int) {
    for i := range naturals {
        squares <- i * i
    }
    close(squares)
}

func printer(squares <-chan int) {
    for v := range squares {
        fmt.Println(v)
    }
}

func main() {
    naturals := make(chan int)
    squares := make(chan int)
    go counter(naturals)
    go squarer(squares, naturals)
    printer(squares)
}
```

### 带缓存的channel

> make(chan T, cap)，cap 大于0就被成为带缓存的channel

### sync.WaitGroup

```go
    wg := sync.WaitGroup{}
    naturals := make(chan int)
    for i := 0; i < 12; i++ {
        wg.Add(1)
        go func(i int) {
            naturals <- i
        }(i)
    }

    go func() {
        for n := range naturals {
            fmt.Println(n)
            wg.Done()
        }
    }()

    wg.Wait()
    close(naturals)
```

### time.ticker

```go
// 只要当整个声明周期都需时才使用
timer := time.Tick(time.Second)

// 可以显式控制，避免发生goroutine泄露
newTimer := time.NewTicker(time.Second)
<-newTimer.C
newTimer.Stop()
```

end
