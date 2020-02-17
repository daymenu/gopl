# 复合数据类型

## 数组

> 数组是由固定长度的特定类型元素所组成的序列

- 数组定义
数组的如果之声明没有赋值，数组中的每个元素会被该类型的零值初始化。数组的长度必须在编译阶段确定。

```go
// 声明数组
var a [3]int
for i,v:=range a{
    fmt.Println(i,v)
}

// 数组个数用...，根据数组元素个数来计算数组个数
var a [...]int{1,2,3}

// 指定索引赋值，会以最大的索引值确定数组长度，其他未显式赋值的会被初始化为对应的零值
var a [...]int{99:1}
```

- 获取数组的长度

```go
var a [3]int
fmt.Println(len(a))
```

## Slice

> 切片代表边长的序列，序列中的每个元素具有相同的类型，只是没有固定的长度，切片引用一个数组对象。len和cap分别返回切片的长度与容量

```go
var sl []int
sl := make([]int, 3)
fmt.Println(sl)

```

## map

> map 也叫哈希表，是一个无序的key/value集合,一个map表就是一个哈希表的引用。

```go
// 字面量语法创建
m := map[int]int{1: 1, 2: 2}

// make 方法创建
m1 := make(map[int]int, 3, 3)

// 这样只是声明没有赋值
var m2 map[int]int
m2[1] = 1 //panic: assignment to entry in nil map

// 更新map中某元素的值
m[1] = 11
// 删除map中的元素
delete(m,1)
```

map中的数据是随机的，所以要依照key的顺序输出，需要先对key进行排序，然后使用key来访问map中的值

## 结构体


end
