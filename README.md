# Golang 学习笔记

# 包管理

## 包

- 一个包下所有 go 文件只能有一个包名
- 包名规则：简洁清晰、全小写
- 入口：包名 main，方法 main，go build 生成二进制文件

## 导入

- 关键字 import
- 本地依赖：Golang 遍历环境变量 GOPATH 中的每一个值，在 $GOPATH/src 寻找本地依赖。可以使用 export 增加 GOPATH，或者修改 ~/.bash_profile，再 source 该文件
- 远程依赖：类似 github.com/money， go get 获取
- 重命名：import myfmt "fmt"
- 禁止导入不使用，会编译错误，需要初始化使用 import _"db/driver/mysql"，会执行 init 函数

### 经典路径

```
- project
-- src
--- package1
---- code1.go
---- code2.go
--- package2 
---- code3.go 
```

code1.go:
```go
package package1

import (
  "package2"
)

func init() {
  package2.DoSomething()
}
```

code3.go:
```
package package2

func DoSomething() {
}

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

## 搭建一个 web server

- 代码在 tserver 中
- 实现了 http 接口功能
- MySQL 操作数据库 CRUD

未实现：
- rest 风格
- ORM 框架
- 日志

### 问题

- 使用 Prepare，select * from some_table where id in(?)，?=1,2，查询结果只有一条，可能原因是防止 SQL 注入产生的问题
