# 函数

## 函数声明

> 函数声明包括函数名、形式参数列表、返回值列表及函数体

## 错误

### 错误处理策略

1. 向上传递错误（类库中一般都将错误向上传递
2. 输出错误
3. 重试
4. 忽略
5. 直接退出（通常在主函数中才直接退出）

## 函数值

> 函数可以像其他变量一样被赋值

```go
func main() {
    pp := printPeople
    pp()
}

func printPeople() {
    fmt.Println("liming")
}
```

## 匿名函数

```go
// defer　后面就是一个匿名函数
defer func(){
    fmt.Println("释放资源")
}()

for i := 0; i < 10; i++ {
    defer func() {
        fmt.Println(i)
    }()
}
// 输出的结果全都是10

for i := 0; i < 10; i++ {
    defer func(i int) {
        fmt.Println(i)
    }(i)
}
// 输出结果是从９～０
```

## defer

> 函数不管以任何形式结束都会调用defer声明的函数，defer 会按照声明的顺序从后向前执行。

## Panic异常

> 程序在运行时才能检测到的错误可以使用panic来抛出异常

## recover 恢复异常

> 在程序运行时如果发生异常可以捕获异常然后做相应的处理

```go
defer func() {
    if p := recover(); p != nil {
        fmt.Println("i catch you", p)
    }
}()
panic("error: this is a important error")
```

end
