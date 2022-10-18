## Go 接口

### 与其他主要编程语言的差异  

1. 接口为非入侵性，实现不依赖于接口定义
2. 所以接口的定义可以包含在接口使用者包内

### 空接口与断言

1. 空接口可以表示任何类型
2. 通过断言来将空接口转换为指定类型

   ``` v,ok:=p.(int) // ok = true``` 时转换成功 

### Go 接口最佳实践

1. 倾向于使用小的接口定义，很多接口只包含一个方法 
   
```
type Reader interface {
	Read(p []byte) (n int, err error)
}
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

2. 较大的接口定义，可以由多个小接口定义组合而成

```
type ReadWriter interface {
   Reader
   Writer
}
```

3. 只依赖于必要功能的最小接口

```
func StoreData(reader Reader) error{
   ...
}
```


