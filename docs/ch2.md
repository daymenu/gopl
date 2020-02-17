# 程序结构

## 命名

1. 命名规则

必须以一个字母（unicode或下划线开头，后面可以跟任意数量的字母、数字或下划线。
2. 关键字

```go
break default func interface select case defer go map struct chan else goto package switch const fallthrough if range type continue for import return var
```

该25个关键字，不能在自定义名字中使用，只能在特定的语法结构中使用
3. 预定义名字

```go
// 内建常量
true false iota nil
// 内建类型
int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr float32 float64 complex128 complex64 bool byte rune string error
// 内建函数
make len cap new append copy close delete complex real imag panic recover
```

预定义关键字可以在自定义名字中使用，但是要慎重，尽量不要使用

## 声明

> 声明语句定义了程序的各种实体对象以及部分或全部属性。go语言主要有四种类型的声明语句：var、const、type和func，分别对应变量、常量、类型和函数实体对象的声明。

```go
package mygo

// Name 我的名字,我的名字可以随便用
var Name = "go"

// money 我的钱，你们谁也别想知道我有几毛钱
var money = 23

// HiName 打招呼
func HiName() string{
    hi:= "hi,"
    return hi + Name
}

// getMoney 看看自己有多少钱
func getMoney() int {
    return Money
}
```

在上面mygo.go的文件中，Name、money属于包一级范围内声明语句，整个包都可以访问，Name还有个特殊之处，其他包也可以访问。首字母大写为导出变量，其他包可以直接访问如mygo.Name

## 变量

> 变量就像我们数学中的x,y一样，它的值是可以变化的。

变量定义形式

```go
// 变量声明
var 变量名 类型 = 表达式

// 简短变量声明
变量名 := 表达式
```

类型或者表达式可以省略其中之一，如果省略类型将会使用类型推断判断该变量是什么类型，省略表达式将会使用零值初始化机制。简短变量声明不能在包以及声明中使用

```go
var name string = "go"

var sex = "man"

var money int

var i, j ,k int

func pi() {
    pi := 3.1415
}

```

零值初始化表
|类型|值|
|:---:|:----:|
|string|空字符串|
|int|0|
|bool|false|
|float64|0|
|complex64|(0+0i)|
|array|[]|

### 指针

一个变量对应一个保存了变量对应类型的内存空间。普通变量在声明语句创建时被绑定到一个变量名，比如叫x的变量，但是还有许多变量始终以表达式的方式引入，例如x[i]或x.f变量。一个指针的值是另一个变量的地址。一个指针对应变量在内存中的存储位置

```go
var x int

var p *int = &x
```

"var x int" 声明语句声明一个x变量，&x表达式将产生一个指向该整数变量的指针，指针对应的类型为*int,指针被称为指向int类型的指针。p为指针名字，那么我们说 “p指向了变量x”或者说 “p保存了x变量的地址” *p 表示 "指针p指向变量的值"

### new 函数

new(T) 将创建一个T类型的匿名变量，初始化为T类型的零值，然后返回变量地址，返回的指针类型为*T

```go
p := new(int)
fmt.Println(*p)

*p = 2
fmt.Println(*p)
```

### 变量的声明周期

变量的声明周期指的是在程序运行期间变量有效存在的时间间隔。对于包一级变量来说，它们的生命周期和整个程序声明周期是一致的。而相比之下，在局部变量的生命周期是动态的，从每次创建一个新变量的声明语句开始，知道该变量不在引用为止，然后变量的存储空间可能被回收。函数参数与返回值变量都是局部变量。

## 赋值

> 使用赋值语句可以更新一个变量的值，最简单的赋值语句是将要被赋值的变量放在=的左边，新值表达式放在=右边

```go
x = 1
*p = true
```

### 元组赋值

元组赋值是另一种形式的赋值语句，它允许同时更新多个变量的值。

```go
x,y = y,x
```

### 可赋值性

赋值语句是显式的赋值形式，但是程序中还有很多地方会发生隐式的赋值行为：函数调用会隐式地将调用参数的值赋值给参数的变量，一个返回语句将隐式地将返回操作的值赋值给结果变量，一个复合类型的字面量也会产生赋值行为。

## 类型

> 变量或表达式的类型定义了对应存储值的属性特征，例如数值在内存的存储大小，它们在内部是如何表达的，是否支持一些操作符，以及它们关联的方法集。

```go
// type 类型名字 底层类型

type money float64
```

money 类型与float64 是两个不同的类型，如果两者参与运算需要类型转化。

## 包和文件

go 语言中的包和其他语言的库或模块的概念类似，目的是为了支持模块化、封装、单独编译和代码重用。

包的声明

```go
package hello
```

包的导入

```go
// 基本导入
import xxx.com/xx/hello
// 导入并重命名
import h xxx.com/xx/hello
// 导入后并不直接使用，一般需要执行一个包的init函数时需要这么做
import _ xxx.com/xx/hello
```

## 作用域

> 一个声明语句将程序中的实体和一个名字关联，比如一个函数或一个变量。声明语句的作用域是指源代码中可以有效使用这个名字的范围。

**声明语句的作用域对应一个源代码的文本区域，它是一个编译时的属性。一个变量的声明周期是指程序运行时变量存在的有效时间段，是一个运行时的概念。**
end
