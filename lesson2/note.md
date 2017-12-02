## Go课程第二节
> Go很多库代码 风格很unix<br>
自举 goruntime 只有go代码没有c

### go通过package组织
- package关键字
- 放在程序的第一行
- 两种package，一种是库package，一种是二进制package
- 二进制package使用main来表示，库package的名字跟go文件所在的目录名一样
> 以`package main`作为文件的第一行<br>
有且只有一个main函数，如echo.go


### go install
> 有利于多个文件的组织<br>
分为两步：<br>
    第一步 go build<br>
    第二步 将执行文件诺至 $GOPATH/bin

### go run 和 go install go build的区别
> go run 针对单个文件<br>
go install 和 go build 针对package

### no install location location for directory GOPATH
> 当前代码没有放在gopath下的src目录<br>
> 如果其他人引用则为以src为根路径的package的全路径

### 命令行参数
> Args命令行数组

### 短变量
> i:=0<br>
s:="hello"<br>
i, j = 0, 1<br>
- 不能放在全局变量里面

### 指针
- *T 即为类型T的指针
- &t 即为去变量t地址
- *p 即为取指针变量所指向的内容

### 网页搜索 
> 广度优先<br>

### gotty
> 远程登陆本机 websocket<br>
tty只是一个终端模拟器