# brew

## brew-文档参考

- 官网：https://brew.sh/



## brew-安装

### mac安装
```bash
# 1.执行脚本
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

# 2.写入path .bashrc = /opt/homebrew/bin
export PATH=/opt/homebrew/bin:$PATH
source .bashrc

# 添加进.zlogin
if [ -f ~/.bashrc ]; then
  source ~/.bashrc
fi
source .zlogin
```

### 安装目录
```bash
# 
cd /usr/local/Cellar

```