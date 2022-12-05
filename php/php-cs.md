# php-cs

## php-cs-安装教程

### ubuntu安装

- php_cs https://www.cnblogs.com/zgaspnet/p/13678180.html

```bash

wget https://cs.symfony.com/download/php-cs-fixer-v2.phar -O php-cs-fixer
chmod +x php-cs-fixer
mv php-cs-fixer /usr/local/bin/
# Tools->External Tools
Program：/usr/local/bin/php-cs-fixer
Arguments：fix "$FileDir$/$FileName$" --config=.php_cs --using-cache=no
Working directory：$ProjectFileDir$

/usr/local/bin/php-cs-fixer fix /xx/xx.php --config=.php_cs --using-cache=no
```