### 程序实体


```go
// 类型别名
type MyString = string
// 除了在名称上 他们是完全相同的
// eg
type byte = uint8
type rune = int32
// 为了代码重构而存在
```

```go
// 类型再定义
type MyString2 string
// string 就是 MyString2 的潜在类型
```
