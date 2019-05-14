# io 

# 继承与组合

Go 的接口是很小的，比如 Reader 和 Writer。想要一个 ReaderWriter，通过内嵌类型（组合模式）来解决代码复用问题，而不是继承模式。

继承模式是 OO 中的核心概念，Go 不是一门 OO 语言，不存在继承。继承是为了多态，组合 + 策略模式也可以实现多态。继承层次过多会使得调用关系变得很复杂，难以理解。

所以在 Go 中屏蔽了继承概念，可以参考 [Why is there no type inheritance?](https://Golang.org/doc/faq#inheritance)

```Gotemplate
// 简单接口
type Reader interface {
	Read(p []byte) (n int, err error)
}

// 简单接口
type Writer interface {
	Write(p []byte) (n int, err error)
}

// 组合接口
type ReadWriter interface {
	Reader
	Writer
}
```

# 数据和方法的正交性

和 OO 语言思考问题方式不一样，Go 强迫程序员保持数据和方法的正交性。数据定义在 struct 中，方法定义在 interface 中。

OO 定义一个学生，学生有姓名字段，学生有读书的的需求

Go 定义一个学生，学生只有姓名字段。读书就是读书的行为。如果学生读书，那么在学生上定义读书方法，正好和读书行为关联上。

# 错误处理

## 返回错误的方式

1. 定义错误，错误是什么
2. 判断出现错误的情况，如果出现错误，直接返回已有的错误，而不是在出错的地方 errors.New，这是一个单例模式，节约资源

## 错误处理

原则是必须处理错误
- 处理不了，return 给上一层
- 到达了最上层，panic 或者返回相应的结果给客户端

