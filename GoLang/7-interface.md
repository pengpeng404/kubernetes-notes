### interface


struct实现接口

1、两个方法的签名完全一致

2、两个方法的名称完全一致



```go
dog := Dog{name : "little dog"}
var pet Pet = &dog
```

Pet 是个 interface，
那么 pet的静态类型就是Pet 接口，
而动态类型就是 *Dog，
pet还可以被其他实现了 Pet接口的对象赋值，
也就是说，
他的动态类型随着赋给他的动态值而变化







































