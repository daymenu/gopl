# 方法

> 将一个特殊的类型与特定的函数关联起来，就是方法

## 方法的声明

```go
type People struct {
    name string
    age int
    sex int
}

// SayHello 自我介绍的方法
func (p People) SayHello() {
    fmt.Printf("Hi,my name is %s and i am %d year old\n", p.name, p.age)
}

```

## 基于指针类型的方法

```go
// ChangeName 自我介绍的方法
func (p *People) ChangeName(name string) {
    p.name = name
}
```

## 结构体嵌套

```go
type Emplyee struct{
    People
    title string // 职称
}

// sayHello 自我介绍的方法
func (e *Emplyee) sayHello() {
    fmt.Printf("Hi,my name is %s and i am %d year old,my title is %s\n", e.name, e.age, e.title)
}

e := Emplyee{People{name: "lili"}, "总经理"}
e.ChangeName("lili")//调用的People类型的方法
e.sayHello()
```

end
