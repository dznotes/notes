# consul-安装教程

## consul-安装

### window安装

```bash
# 1、下载：https://www.consul.io/downloads
# 2、解压：
# 3、启动：启动consul agent
# 4、测试：访问 http://localhost:8500/

# 可以把consul所在路径添加进环境变量。

# 默认方式1 使用localhoast
./consul.exe agent -dev

# 方式2 指定本地IP
./consul.exe agent -dev -client 192.168.xx.x

# 方式3 window10开机启动
sc.exe create "Consul_1.9.3" binPath="C:\shell\consul.exe agent -dev"
sc.exe delete "Consul_1.9.3" # 删除服务
sc.exe start "Consul_1.9.3" # 启动服务
```
