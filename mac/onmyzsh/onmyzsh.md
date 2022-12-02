# onmyzsh

## onmyzsh-安装

### mac安装

```bash
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
```

## onmyzsh 插件安装

```bash
cd
cd .oh-my-zsh/plugins
git clone https://github.com/zsh-users/zsh-autosuggestions.git
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git



# 添加插件
cd
vim .zshrc
plugins=(git zsh-syntax-highlighting zsh-autosuggestions)
```