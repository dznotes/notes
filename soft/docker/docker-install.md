# docker-安装教程

## docker-安装

### mac安装

### mac M1安装

- https://docs.docker.com/desktop/mac/apple-silicon/

### Ubuntu安装

```bash
# 更新ap包
sudo apt-get update

# 安装
curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun

# 测试
docker run hello-world
```

## docker-compose-安装

### ubuntu安装

```bash
# 查看版本：https://github.com/docker/compose/releases

# 下载
curl -L https://github.com/docker/compose/releases/download/1.29.2/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose 
docker-compose --version
```