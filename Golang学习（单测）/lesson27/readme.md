## Benchmark

```go
func BenchmarkConcatStringByAdd(b *testing.B){
	// 与性能测试无关的代码
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		// 测试代码
    }
	b.StopTimer()
	// 与性能测试无关的代码
}
```
