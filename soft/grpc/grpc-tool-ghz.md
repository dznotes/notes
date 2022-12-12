# ghz

- 基准测试和负载测试工具

## ghz-文档参考

- 官网：https://ghz.sh/
- github：https://github.com/bojand/ghz 
- 教程：
    - https://lixueduan.com/post/grpc/10-benchmark/

## ghz-安装

```bash
git clone https://github.com/bojand/ghz
cd ghz

make build

cd cmd/ghz
go build 

#【可选】
mv ghz /usr/local/bin/
ghz -version
```

## ghz-使用

```bash
# 测试
ghz --insecure  \
--proto ./hello.proto \
--call Service.Ping \
-d '{"request": "1111"}'  \
-n 2000 \
-c 20 \
127.0.0.1:8080

```
## ghz-参数

```bash
-rps	每秒请求数 (RPS) 
-c	并发数
-n	总请求数
```