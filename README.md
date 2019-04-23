# Golang 学习笔记

## 依赖管理

- Golang 遍历环境变量 GOPATH 中的每一个值，在 $GOPATH/src 寻找依赖
- 可以使用 export 增加 GOPATH，或者修改 ~/.bash_profile，再 source 该文件
- go get -t 远程获取依赖

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

## 常用命令

- go vet 检查常见代码错误
- go install 安装依赖
- go build 编译产生可执行文件
- go fmt 格式化代码

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
