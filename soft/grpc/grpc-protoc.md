# protoc

## protoc-文档参考


## protoc-安装教程


- 下载地址：https://github.com/protocolbuffers/protobuf/releases

### 安装步骤

- 选择对应的版本下载，解压后移动到可执行目录。
- 移动到$GOPATH/bin目录下或者/usr/local/bin。

```bash
# Ubuntu
mv protoc /usr/local/bin/

# 测试
protoc --version
libprotoc 3.17.1
```

## protoc-使用


- -I / --proto_path：指定源proto文件和依赖的proto文件的目录路径
- 指定希望编译的proto文件路径
- 指定生成的代码要存放的目标目录
```bash
protoc -I xx ABC/DEF/proto --go_out=plugins=grpc:<xxxx>/xxx

protoc -I xx \
ABC/DEF/proto \
--go_out=plugins=grpc:<xxxx>/xxx
```