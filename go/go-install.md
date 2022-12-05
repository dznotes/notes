# golang-安装教程

## golang-文档参考


## golang-安装

- 官网下载：https://golang.org/dl/
- https://www.cnblogs.com/00986014w/p/10001093.html
- https://cloud.tencent.com/developer/article/1623121
### mac安装

```bash
# M1芯片安装不了Go的14、15版本
# 直接下载：https://golang.org/dl/go1.16.5.darwin-arm64.pkg

# 添加可执行文件到bashrc
cd
vim .bashrc
export PATH=/usr/local/go/bin:$PATH
export PATH=/Users/xx/shell:$PATH

# 退出，保存
source .bashrc
```

### Ubuntu安装

```bash
# 方式1
# 下载
wget https://dl.google.com/go/go1.14.linux-amd64.tar.gz
# 解压
tar zxvf go1.14.linux-amd64.tar.gz -C /usr/local

# 方式2
# 安装在/usr/local/go目录下
# wget -c https://dl.google.com/go/go1.14.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local
```