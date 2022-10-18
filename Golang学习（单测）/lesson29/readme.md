## reflect.TypeOf vs reflect.ValueOf

* reflect.TypeOf 返回类型 (reflect.Type)
* reflect.ValueOf 返回值 (reflect.Value)
* 可以从 reflect.Value 获得类型
* 通过 kind 来判断类型

## 判断类型 - Kind()

```go
const(
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
...	
)
```