# 接口

## 接口定义

> 接口是一种类型,接口类型是一种抽象的数据类型。通过接口定义可以做什么事情。

```go
type Speaker interface {
    speak()
}
```

## 实现接口的条件

一个类型如果实现了一个接口的所有方法，我们就说这个类型实现了这个接口

## 接口值

一个接口值分为具体类型与对应类型的那个值，一个包含nil指针的接口不是nil接口

## 类型断言

类型断言是一个使用在接口值上的操作。语法上看起来像x.(T)被称为断言类型，这里x表示一个接口的类型和T表示一个类型。
1. 如果断言的类型是一个具体类型，然后类型断言检查x的动态类型是否和T相同。如果这个检查成功了，类型断言的结果是x的动态值，当然它的类型是T。
2. 如果断言的类型T是一个接口类型，然后类型断言检查是否x的动态类型满足T。如果检查成功，动态值没有获取到，这个结果仍然是一个具有相同类型和值部分的接口值，但是结果友类型T。对一个接口类型的类型断言改变了接口的表述方式，改变了可以获取的方法集合

## 类型开关

```go
// a.(type) 会返回a的类型
var a interface{}
a = "aa"
switch a.(type) {
case int:
    fmt.Println(" i am int")
case string:
    fmt.Println("i am string")
}
```

## time 使用

```go
nowTime := time.Now()
t := nowTime.Format(ShowTime)

tenAfter := nowTime.Add(10 * time.Minute)

d := tenAfter.Sub(nowTime)

fmt.Println(d)

fmt.Println(t)

fmt.Println(tenAfter.Format(ShowTime))

fmt.Println(time.Parse(ShowTime, tenAfter.Format(ShowTime)))
```

end
