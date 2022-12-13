# go-mod

## go-mod-使用教程

### 开启GO111MODULE

- on：支持go mod模式
- off：不支持go mod模式
- auto (默认模式)：
  - 如果代码在gopath下，则自动使用gopath模式；
  - 如果代码不在gopath下，则自动使用go mod模式。


```bash
# 设置环境变量
go env -w GO111MODULE=on
go env -w GO111MODULE=off
go env -w GO111MODULE=auto
```

### go-proxy

- https://goproxy.cn
- https://goproxy.io
- https://proxy.golang.org

```bash

$ go env -w GOPROXY=https://goproxy.cn,direct
https://goproxy.cn,https://goproxy.io,direct
```


## go-mod-常用命令



| 命令                                    | 含义                                                         |
| :-------------------------------------- | :----------------------------------------------------------- |
| go mod init xx                          | 初始化xx项目（项目开始的时候使用）                           |
| go mod download xx                      | 下载modules到本地cache：$GOPATH/pkg/mod和 $GOPATH/pkg/sum 下 |
| go mod tidy                             | 删除不需要的依赖包，新增需要的依赖包                         |
| go mod vendor                           | 拉取需要的依赖包，生成vendor目录                             |
| go mod graph                            | 把模块之间的依赖图显示出来                                   |
| go mod verify                           | 校验依赖                                                     |
| go mod why                              | 解释为什么需要依赖                                           |
| go mod edit                             | 编辑go.mod文件                                               |
| go mod edit -fmt go.mod                 | 格式化go.mod                                                 |
| go mod edit -replace=xx@v1.1.=yy@latest | 替换包源（使用后者替换前者）                                 |
| **其他相关命令**                        |                                                              |
| go clean -modcache                      | 清理本地cache                                                |
| go get -u xx                            | 更新依赖到最新版本                                           |
| go get -u xx@v1.xxx.yy                  | 更新到指定版本                                               |
| go help mod                             | 查看mod帮助                                                  |



```bash
go help mod

# 输出
Go mod provides access to operations on modules.

Note that support for modules is built into all the go commands,
not just 'go mod'. For example, day-to-day adding, removing, upgrading,
and downgrading of dependencies should be done using 'go get'.
See 'go help modules' for an overview of module functionality.

Usage:

	go mod <command> [arguments]

The commands are:

	download    download modules to local cache
	edit        edit go.mod from tools or scripts
	graph       print module requirement graph
	init        initialize new module in current directory
	tidy        add missing and remove unused modules
	vendor      make vendored copy of dependencies
	verify      verify dependencies have expected content
	why         explain why packages or modules are needed

Use "go help mod <command>" for more information about a command.
```