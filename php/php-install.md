# php-安装教程

## php-安装

### mac安装

```bash
brew search php
brew install php@8.0

echo 'export PATH="/opt/homebrew/opt/php@8.0/bin:$PATH"' >> ~/.zshrc
echo 'export PATH="/opt/homebrew/opt/php@8.0/sbin:$PATH"' >> ~/.zshrc
source .zshrc
php --version

# 方法2
brew tap shivammathur/php
brew search php
brew install shivammathur/php/php@7.0
```