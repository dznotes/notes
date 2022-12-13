# golang

## golang-文档参考

- github官网：https://github.com/golang/go
- https://github.com/avelino/awesome-go
- go夜读：https://github.com/talkgo/read

### 教程/书籍

-《Go 编程基础》-无闻：https://github.com/unknwon/go-fundamental-programming
- 雨痕学习笔记：https://github.com/qyuhen/book
- 收集网络上的开发书籍（golang）：https://github.com/zhizouxiao/dev-books

### 框架

- beego：https://github.com/beego/beego

### gopher大会
- 会议PPT：https://github.com/gopherchina/conference

### 开发工具
- json=>go：https://mholt.github.io/json-to-go/   
- https://github.com/mholt/json-to-go

## 三方包

- errors with stacktraces for go：https://github.com/go-errors/errors
- a cron library for go：https://github.com/robfig/cron
- Cron expression parser in Go language (golang)：https://github.com/gorhill/cronexpr
- Visualize call graph of a Go program using Graphviz：https://github.com/ofabry/go-callvis

### 日志

## go-问题解决

### golang.org/x/..包导入问题

- 解决思路：golang在github上建立了一个镜像库，下载github上的镜像库放入GOPATH下即可

```bash
mkdir -p $GOPATH/src/golang.org/x
cd $GOPATH/src/golang.org/x
git clone https://github.com/golang/sync.git
git clone https://github.com/golang/crypto.git
git clone https://github.com/golang/sys.git
git clone https://github.com/golang/net.git 
git clone https://github.com/golang/time.git
git clone https://github.com/golang/text.git
git clone https://github.com/golang/term.git
git clone https://github.com/golang/oauth2.git

cd $GOPATH/src/google.golang.org
git clone https://github.com/protocolbuffers/protobuf-go.git
mv protobuf-go protobuf
```