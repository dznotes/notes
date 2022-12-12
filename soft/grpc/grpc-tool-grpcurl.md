# grpcurl

## grpcurl-文档参考

- github：https://github.com/fullstorydev/grpcurl 

## grpcurl-安装教程

```bash
# 安装
go get github.com/fullstorydev/grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl

# 
grpcurl -version
```

## grpcurl-使用教程

```bash
grpcurl 127.0.0.1:12345 list

# 错误解决
Failed to dial target host "127.0.0.1:12345": tls: first record does not look like a TLS handshake

# 使用 -plaintext
# 列出服务名称
grpcurl -plaintext 127.0.0.1:12345 list
 

# 列出方法名称列表
grpcurl -plaintext 127.0.0.1:12345 list serviceName


# 列出服务定义
grpcurl -plaintext 127.0.0.1: 12345 describe serviceName

# 列出方法定义
grpcurl -plaintext 127.0.0.1: 12345 describe serviceName.methodName

# 列出类型定义
# serviceName去除当前服务名称-1个级别
grpcurl -plaintext 127.0.0.1: 12345 describe serviceName.xxRequest

# 调用

# 调用方法
grpcurl -d @ -plaintext 127.0.0.1: 12345 serviceName.methodName << EOM
{
  "request": "1111"
}
EOM
```