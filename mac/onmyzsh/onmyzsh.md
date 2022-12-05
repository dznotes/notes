# onmyzsh

## onmyzsh-文档参考

- 官网：https://ohmyz.sh/

## onmyzsh-安装

### mac安装

```bash
sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
```
## onmyzsh-常用插件-官方

- 插件列表：https://github.com/ohmyzsh/ohmyzsh/tree/master/plugins

| 插件   | 解释                 |
| ------ | -------------------- |
| git    | 【常用】git 命令别名 |
| golang | 【常用】go命令别名   |
|        |                      |

## onmyzsh-常用插件-社区

- https://github.com/orgs/zsh-users/repositories?type=all

| 插件                    | 解释     |
| ----------------------- | -------- |
| zsh-autosuggestions     |          |
| zsh-syntax-highlighting | 语法高亮 |

## onmyzsh-插件安装

```bash
cd ~/.oh-my-zsh/custom/plugins
git clone https://github.com/zsh-users/zsh-autosuggestions.git
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git

# 添加需要的插件
vim ~/.zshrc
plugins=(
	git 
	zsh-syntax-highlighting 
	zsh-autosuggestions
)
```

