## Go 的错误机制

### 与其他主要编程语言的差异

1. 没有异常机制
2. error 类型实现 error 接口
```
type error interface{
    Error() String
}
```
3. 可以通过 errors.New 来快速创建错误实例
```
errors.New("n must be in the range [0,1]")
```

### 返回错误建议

建议尽早返回错误信息，避免多层嵌套

即 ``` if err!=nil{}```的写法


### panic

* panic 用于不可以恢复的内容
* panic 退出前会执行 defer 指定的内容

### panic vs os.Exit

* os.Exit 退出时不会调用 defer 指定的函数
* os.Exit 退出时不输出当前调用栈信息