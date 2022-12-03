# etcd-安装教程

## etcd-安装

### mac 安装


```bash
# 安装
brew install etcd

# 启动1:直接启动
/opt/homebrew/opt/etcd/bin/etcd

# 启动2:后台启动
brew services start etcd

# 启动3:直接启动
etcd

# 端口被占用
sudo lsof -i :2380
kill -9 pid
```

### ubuntu安装

```bash
# 安装etcd
sudo apt-get install etcd-server
# 重启
service etcd restart
# 输出
$ etcd -version
etcd Version: 3.2.17
Git SHA: Not provided (use ./build instead of go build)
Go Version: go1.10
Go OS/Arch: linux/amd64

# 安装etcdctl
sudo apt-get install etcd-client

# 输出版本
$ etcdctl -v
etcdctl version: 3.2.17
API version: 2

# 重要：改为使用API version = 3
vim .bashrc 
export ETCDCTL_API=3
source .bashrc
```