# gitbook

## gitbook-文档参考
- [Mac中安装和使用GitBook](https://xiaochaowei.com/2020/01/10/InstallGitBookOnMAC/)
- [gitbook 发布 github pages](https://blog.csdn.net/xixihahalelehehe/article/details/125115061)

## gitbook-安装

### Mac安装

- 使用安装包安装node.js(官网: https://nodejs.org/en/)

```bash
# 查看版本号
node -v
```
- 安装gitbook

```bash
- 安装gitbook
sudo npm install gitbook-cli -g

# 查看版本号

- 初始化项目
gitbook -V

```
- 使用gitbook

```bash
# 初始化
gitbook init

# 运行项目
gitbook serve

# 本地访问
http://localhost:4000
```

## gitbook-常用命令

| 命令            | 含义                              |
| :-------------- | :-------------------------------- |
| gitbook -V      | 查看版本                          |
| gitbook init    | 初始化gitbook                     |
| gitbook serve   | 本地运行gitbook服务，默认端口4000 |
| gitbook install | 插件安装                          |
| gitbook build   | 生成html                          |

## gitbook-常用插件

| 插件 | 说明                  |
| :--- | :--------------------- |
| back-to-top-button | 返回顶部 |
| chapter-fold |  |
| code | 代码添加行号&复制按钮 |
| copy-code-button | 代码复制按钮 |
| search | 默认搜索（不支持中文） |
| search-pro | 高级搜索（支持中文，关闭 -lunr、-search） |
| splitter | 侧边栏宽度调整 |
| tbfed-pagefooter | 页面添加页脚 |

### gitbook-常见主题

| 插件           | 说明     |
| :------------- | :------- |
| theme-default  | 默认主题 |
| theme-comscore | 彩色标题 |
