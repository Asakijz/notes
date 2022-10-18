## 字符串

### 与其他主要编程语言的差异
1. string 是数据类型，不是引用或指针类型
2. string 是只读 byte slice，len 函数可以取它所包含的 byte 数
3. string 的 byte 数组可以存放任何数据

## Unicode UTF-8
1. Unicode 是一种字符集（code point）
2. UTF-8 是 unicode 的存储实现（转换为字节序列的规则）

## 编码与存储
第一行是“中”字在这个字符集里的编码

第二行是“中”物理存储在这个 UTF-8 规则

第三行是“中”存放在 byte 的这个切片中，每个 byte 放一个

| 字符            | “中”              |
|---------------|------------------|
| Unicode       | 0x4E2D           |
| UTF-8         | 0xE4B8AD         |
| string/[]byte | [0xE4,0xB8,0xAD] |