# git

## git-文档参考

- git官网：https://git-scm.com/
- git window 中文镜像下载： https://npm.taobao.org/mirrors/git-for-windows/
- git常用命令速查表：https://www.w3cschool.cn/git/git-cheat-sheet.html


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

# 添加
ssh-add ~/.ssh/id_rsa

# 验证
ssh -T git@github.com
```

## git-常用命令



| 命令                                    | 含义                                              |
| :--------------------------------------- | :------------------------------------------------ |
| **创建**                                |                                                   |
| git init                                | 初始化本地git仓库，使用当前目录作为git仓库        |
| git init xx_name                        | 使用指定名称作为git仓库                           |
| git clone xx.git                        | 克隆远程仓库                                      |
| **提交历史记录**                        |                                                   |
| git log                                 | 显示日志                                          |
| git show commit_id                      | 显示某个提交的详细内容                            |
| git blame xx_filename                   | 在每一行显示commit号                              |
| **本地更改**                            |                                                   |
| git status                              | 查看当前版本状态+是否修改                         |
| git diff                                | 查看所有变更                                      |
| git diff xx_file_name                   | 查看具体文件变更                                  |
| git add x.txt                           | 添加单一文件到缓存                                |
| git add --all                           | 添加全部文件到缓存                                |
| git commit -m "commit info"             | 提交                                              |
| git commit -am "commit info"            | 将add和commit合为一步                             |
| **分支和标签**                          |                                                   |
| git branch                              | 显示本地分支                                      |
| git branch -v                           | 查看分支-带有最后提交信息                         |
| git branch -d xx_branch_name            | 删除分支                                          |
| git checkout xx_branch_name             | 切换到指定分支                                    |
| git checkout -b xx_branch_name          | 新建并切换到新的分支                              |
| git checkout xx_file_name               | 丢弃本地更改信息                                  |
| git tag xx_tag_name                     | 给当前分支打标签                                  |
| **更新和发布**                          |                                                   |
| git remote -v                           | 查看远程分支详细信息                              |
| git remote show xx_branch_name          | 查看某个分支信息                                  |
| git remote add origin xxx.git           | 添加一个新的远程仓库                              |
| git fetch origin xx_branch_name         | 获取远程分支，但不更新本地分支，另外需要merge操作 |
| git pull origin xx_branch_name          | 获取远程分支，并更新本地分支                      |
| git push origin xx_branch_name          | 推送本地分支到远程分支                            |
| git push origin --delete xx_branch_name | 删除远程分支                                      |
| git push --tags                         | 推送本地标签                                      |
| **撤销**                                |                                                   |
| git reset --hard                        | 将当前版本重置为HEAD                              |
| git reset --hard commit_id              | 将当前版本重置为指定commit_id                     |
| git revert commit_id                    | 撤销提交                                          |
| **合并**                                |                                                   |
| git merge xx_branch_name                | 合并分支到当前分支，2个分支                       |
| git merge --abort                       | 回到执行merge之前                                 |
| git rebase xx_branch_name               | 合并分支到当前分支，1个分支                       |
| git rebase --abort                      | 回到执行rebase之前                                |
| **配置**                                |                                                   |
| git config --list                       | 显示当前的git配置                                 |
| git config --global user.name "xx"      | 设置全局的user.name信息                           |
| git config --global user.email "xx"     | 设置全局的user.email信息                          |



## git-gui

- sourcetree官网：https://www.sourcetreeapp.com/

## git-问题解决

### OpenSSL SSL_connect报错

- 问题描述：OpenSSL SSL_connect: Connection was reset in connection to github.com:443
- 问题解决

```bash
git config --global http.sslVerify false
```