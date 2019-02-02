### 解释为什么帮助信息在它的默认值是20.0没有包含°C的情况下输出了°C?

1. 猜想：猜想默认值的时候他会调用Celsius类型的String()方法
2. 验证：修改如下代码
```go
func (c Celsius) String() string {
	fmt.Printf("%f\n", c)
	return fmt.Sprintf("%.2f °C", c)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	fmt.Println("----pre-----")
	flag.CommandLine.Var(&f, name, usage)
	fmt.Println("----last-----")
	return &f.Celsius
}
```

发现在执行这一行代码的时候，会输出String中自定义的输出,flag.CommandLine.Var(&f, name, usage),说明此行代码调用了Celsius类型的String方法，所以会在默认值的情况下输出°C

flag.CommandLine.Var(&f, name, usage)源码
```go
func (f *FlagSet) Var(value Value, name string, usage string) {
	// Remember the default value as a string; it won't change.
	flag := &Flag{name, usage, value, value.String()}//这里调用了value的String()
	_, alreadythere := f.formal[name]
	if alreadythere {
		var msg string
		if f.name == "" {
			msg = fmt.Sprintf("flag redefined: %s", name)
		} else {
			msg = fmt.Sprintf("%s flag redefined: %s", f.name, name)
		}
		fmt.Fprintln(f.Output(), msg)
		panic(msg) // Happens only if flags are declared with identical names
	}
	if f.formal == nil {
		f.formal = make(map[string]*Flag)
	}
	f.formal[name] = flag
}
```