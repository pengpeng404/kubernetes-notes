### 函数




函数是一等公民 --> 函数式编程



问题1：如何理解闭包：

```go
func genCalculator(op operate) calculateFunc {
	return func(x int, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}
```

以为引用了自由变量，
而呈现出一种不确定的状态，
也就是说函数的内部逻辑是不完整的，
需要这个自由变量参与完成，
比如例子中的 op


这个例子返回了一个匿名函数，
这就是一个闭包函数，
函数中的变量 op，
既不是函数参数，
也不是自己声明的，
是一个自由变量，
这个自由变量代表什么，
并不是定义闭包函数的时候确定的，
而是在调用 genCalculator 时候确定的






















