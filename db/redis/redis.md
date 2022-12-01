# redis

## redis-文档参考

## redis-安装

### mac安装

```bash
# brew 安装
brew install redis

# 启动方式1：后台启动
brew services start redis
 
# 启动方式2:
/opt/homebrew/opt/redis/bin/redis-server /opt/homebrew/etc/redis.conf
```

### Ubuntu安装

- 安装

```bash
# apt-get 安装
sudo apt-get install redis-server
```

- 设置密码

```bash
# 设置密码
sudo vim /etc/redis/redis/redis.conf

# 修改 配置文件里面的 requirepass，把注释删除
# 修改密码

# 重启
service redis restart
```

- 查看状态

```bash
# 检查redis服务器系统进程
ps -aux | grep redis

# 检查redis服务器状态
netstat -nlt | grep 6379
```

