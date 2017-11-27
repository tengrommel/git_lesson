## 为什么要用golang

> 代替c 适合高并发 编译各个平台 包依赖 不需要gcc <br>
基础设置的管理 c++ 用c的性能提供灵活的封装 <br>
程序员学习语言并不是炫技而是为了更好的完成项目的<br>
python是一种表达性的语言 golang的特性：高并发 高性能<br>
golang编译只和内核有关并不和gcc有关 稳定性好 静态编译

### 基本特点

- 静态编译
- 垃圾回收
- 简洁的符号和语法
- 平坦的类型系统
- 基于CSP的并发模型
- 高效简单的工具链
- 丰富的标准库
- 适合分布式开发

### 并发和并行

- 并发(concurrent)不是并行(parallel)
- 一个例子，node.js具有并发的能力，但不能充分利用多核
- 写出一个能充分利用多核的程序需要很深的系统编程积淀
- 得益于优秀的设计，go可以轻松地写出跑满所有CPU的程序

### go语言的应用

- Docker
- kubernetes
- etcd

### 基础知识

- 编译器 go命令
- 开发环境和运行环境不同
- window平台开发下载二进制
> 源码编译 <br>
1.5之前：编译版本下载 go1.4.3 c编译器 ./make.bash (gcc依赖) <br>
1.9之后：自举  需要设计环境变量和golang的编译器<br>
GOROOT<br>
GOPATH<br>
PATH<br>
- 其他平台的编译
> GOOS=linux go build hello.go<br>
GOOS=windows go build hello.go
-window开发编译
> set GOOS=linux<br>
go build hello.go
- 代码规范 go fmt -w hello.go
- 自动引入包 goimports -w hello.go
- 启动 HTTPS_PROXY=http://192.168.1.53:20000 atom 代理

### GOPATH
- src
- bin
- pkg
- go get 下载并编译
- godoc -http=:9090

### 指针
> 什么是变量<br>
内存的标识符<br>
变量里存的不是内容而是另一个变量的地址<br>
c语言中的指针指向不存在对象的指针 称野指针<br>
