# git
## 参考
- 

## git安装
### Mac 安装
```bash
brew install git
```
### Ubuntu安装
```bash
sudo apt-get install git
```

## git-ssh-key
### 生成步骤
```bash
git config --global user.name "xx"
git config --global user.email "xx@xx.com"
git config --list

# 使用rsa生成
ssh-keygen -t rsa -C "邮箱"

# 使用ed25519生成
ssh-keygen -t ed25519 -C "邮箱"

# 不使用邮箱（使用这个）
ssh-keygen -t rsa

# 查看
cat ~/.ssh/id_rsa.pub
```