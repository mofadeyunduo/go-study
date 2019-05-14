# Go 学习笔记

# Git hook

- 利用 git hook 和 go fmt 等命令对 go 程序进行检查 
- 脚本位置 $PROJECT/.git/hooks/pre-commit

# Go 特点

- 开发速度快，语法简单，编译速度快
- 并发性能高，语法方便，goroutine
- 使用组合而不是继承，基本没有继承的概念
- 内存自动管理

# Go 开发工具

- http://play.golang.org
- Goland

# 变量 var

- var name type = {} or 1 or "" or true
- := 声明并赋值
- 大写公开，小写私有
- 初始为各个类型的空值，数字为 0，字符串为 ""，布尔为 false，指针为 nil
- 支持闭包，循环拿到的是值而不是副本，所以在循环中用 goroutine 的时候不能应用循环的值
- 标签语法 `json:"name""`

# 通道 chan

- 打开 make(chan *value)，关闭 close(*vchan)，close 会阻塞
- 存入 vchan <- somevalue, 读取 somevalue <- vchan
- 常用 select 接受值， select <- value {case val: doSomething(); defalut:}
- 数据同步，一个地方存入，一个地方获取
- 可以不带缓冲，也可以带缓冲，make(*vchan, 10) 缓冲大小为 10

# 函数 func

- 允许有多个返回值
- 参数是值传递，修改无效；除非传递的是指针
- 可以声明接受者 receiver，如果想要修改，必须声明指针接受者

# 切片 slice

- []type or make([]type)
- 动态大小的数组
- for index, value := range slice 遍历
- append 增加元素

# 映射 map

- map[key type]value type
- make 生成

# goroutine
 
- go func() 启动

# 常用模式

- 同步 goroutine： sync.WaitGroup 首先定义 var sync.WaitGroup，其次增加次数 wg.Add(2)，之后在需要等待的地方 wg.Wait(),最后在 goroutine 中 wg.Done()
- 检查 error： if err != nill {doSomething();}

# 包管理

## 包

- 一个包下所有 go 文件只能有一个包名
- 包名规则：简洁清晰、全小写
- 入口：包名 main，方法 main，go build 生成二进制文件

## 导入

- 关键字 import
- 本地依赖：首先找 GOROOT；Golang 遍历环境变量 GOPATH 中的每一个值，在 $GOPATH/src 寻找本地依赖。可以使用 export 增加 GOPATH，或者修改 ~/.bash_profile，再 source 该文件
- 远程依赖：类似 github.com/money， go get 获取
- 重命名：import myfmt "fmt"
- 禁止导入不使用，会编译错误，需要初始化使用 import _"db/driver/mysql"，会执行 init 函数

## 一些代码合作相关的建议

- 分享给他人时，包应该在代码库的根目录中
- 包内容不要过多
- go fmt
- 写文档

## 依赖管理

- go module，官方
- 初始化
```gotemplate
go mod init packageName
```

### GoLand 设置

shift shift 搜索 GOPATH 设置环境变量

# Go 常用命令

- go vet 检查常见代码错误，如 printf 参数错误、方法签名错误等等
- go install 安装依赖
- go build 编译产生可执行文件，... 表示目录下所有文件
- go run 编译后执行
- go fmt 格式化代码，可以配合 git 强制 fmt
- go clean 清除可执行文件
- go test 执行测试
- go doc 查看文档，go doc tar 查看 tar 文档，godoc -http=:8080，生成 web 文档。可以在目录下建 doc.go，生成包的注释

### 问题

- 使用 Prepare，select * from some_table where id in(?)，?=1,2，查询结果只有一条，可能原因是防止 SQL 注入产生的问题
