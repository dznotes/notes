# grpcui

## grpcui-安装教程

```bash
export GOPROXY=https://goproxy.cn
go get github.com/fullstorydev/grpcui
go install github.com/fullstorydev/grpcui/cmd/grpcui

# 到go build生成的目录
cd go/bin/
chmod 777 grpcui
sudo mv grpcui /usr/local/bin/

grpcui -help


# 需要修改端口，指定端口即可
grpcui -plaintext 127.0.0.1:8080
```