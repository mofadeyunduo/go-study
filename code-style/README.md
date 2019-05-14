# 风格

## 枚举怎么写

```gotemplate
type referer string

const (
	bpd referer = "bdp"
)
```

## 错误处理，首先全局定义

所有的错误都要全局定义再使用
```gotemplate
type BusinessError struct {
	Code int
	error
}

var (
	NotExist = BusinessError{10001, errors.New("item not found")}
)
```

## return 的写法

- 函数体定义返回变量
- 需要 return 直接 return，不用带返回值

```gotemplate
func test() (i interface{}, err error) {
    j, err := init()
    if err != nil {
        return
    }
    i = operate(j)
    return
}
```

## channel 怎么写避免死锁


