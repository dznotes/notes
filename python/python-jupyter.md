# jupyter

## jupyter-文档参考

- 官网：https://jupyter.org/
- 参考：https://www.cnblogs.com/xiaoqi/p/6393677.html
- https://www.cnblogs.com/faramita2016/p/7512471.html

## jupyter-安装教程

### Ubuntu安装

```bash
# 前提安装pip3

#使用pip3安装 jupyter
pip3 install jupyter

# 添加.bashrc
vim .bashrc
export PATH=$PATH:~/.local/bin
source .bashrc

# 进入指定目录，运行jupyter
mkdir jnote
cd jnote
jupyter notebook

# 访问
http://localhost:8888/

# 可选
# 安装matplotlib包
python3 -m pip install matplotlib 
```