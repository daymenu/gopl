# 基础数据类型

> 所有的数据底层都是由比特组成，一般计算机操作的是固定大小的数，比如整数、浮点数、比特数组、内存地址等。

**go语言的数据类型分类：**
|类型|值|
|:---:|:----:|
|基础类型|数字、字符串、布尔|
|复合类型|数组、结构体|
|引用类型|通道（chan） 切片（slice） 指针 字典（map） 函数|
|接口类型|interface|

## 整型

|类型名称|比特位|取值范围|
|:---:|:----:|:----:|
|int8|8|-128~127|
|int16|16|-32768~32767|
|int32|32|-2147483648 ~ 2147483647|
|int64|64|-9223372036854775808 ~ 9223372036854775807|
|uint8|8|0~255|
|uint16|16|0~ 65535|
|uint32|32|0~ 4294967295|
|uint64|64|0~18446744073709551615|

该值可以通过math.MinInt16等一系列函数获取

### 运算符优先级(二元运算符遵循从左向右结合)

|优先级|
|:---:|
|* / % & &^|
|+ - | ^|
|== != < <= > >=|
|&&|
|\|\||

## 浮点数

go 语言提供两种浮点数 float32和float64,实现的标准是IEEE754

|类型|取值范围|
|:---:|:---:|
|float32|3.4028234663852886e+38|
|float64|1.7976931348623157e+308|

## 复数

go语言提供两种类型的复数类型complex64和complex128,分别对应float32与float64的精度

```go
var x c1 = complex(1,2)
cr1 := real(c1)//获取实数实部
ci1 := imag(c1)//获取实数虚部
```

## 布尔

一个布尔类型的值只有两种 true和false。if与for的条件部分都是bool类型的值。

## 字符串

一个字符串是一个不可改变的字节序列。文本字符串通常被解释为采用utf8编码的Unicode码点序列，utf8是可变的编码一个字符可以占一字节到四字节不等。

```go
s := "hello 中国"
fmt.Println(len(s))// len 返回字符串的字节长度
fmt.Println(utf8.RuneCountInString(s)) // 返回字符串中的utf8码点个数

sRune := []rune(s) // 将字符串转化为rune序列
for i := range sRune {
    fmt.Print(string(sRune[i]))
}

// 遍历utf8字符
for i := 0; i < utf8.RuneCountInString(s); {
    sep, j := utf8.DecodeRuneInString(s[i:])
    fmt.Println(string(sep), j)
    i += j
}
```

## 字符串与数字的相互转换

```go
// 字符串转化为整型
var age int
var int1 int64
var int2 int64

str := "32"
str1 := "922337203685477580"

age, _ = strconv.Atoi(str)
int1, _ = strconv.ParseInt(str1, 10, 32)
fmt.Sscanf(str1, "%d", &int2)

fmt.Println(age)
fmt.Println(int1)
fmt.Println(int2)

// 整型转化到字符串
var age64 int64 = 32
var age int = 32
var ageStr string

ageStr = strconv.FormatInt(age64, 10)
ageStr = fmt.Sprintf("%d", age)
ageStr = strconv.Itoa(age)

fmt.Println(ageStr)
```

## 常量

常量的值编译阶段计算，常量值的基本类型　布尔、字符串、数字，常量是不可以修改的。
const() 括号中可以批量声明，如果某一行没有值，那么它的值会与倒着往上走的第一个的值相等。

### iota 常量生成器

iota 在表达式里以０开始然后在每行递增

```go
const (
//Monday 星期一
Monday = iota
//Tuesday 星期二
Tuesday
//Wednesday 星期三
Wednesday
//Thursday 星期四
Thursday
//Friday 星期五
Friday
//Saturday 星期六
Saturday
//Sunday 星期日
Sunday
)
```

const 中出现iota被重置为０，const中每增加一行常量声明iota增加１
end
