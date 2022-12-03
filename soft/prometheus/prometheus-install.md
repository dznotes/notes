# prometheus-安装教程

## prometheus-安装

### ubuntu安装

```bash
# 访问：https://prometheus.io/download/
# 下载组件prometheus/prometheus：alertmanager-0.23.0.linux-amd64.tar.gz
# 解压到目标文件夹：prometheus
# 到目标文件夹中

# 启动：
./prometheus

# 使用配置文件启动:配置prometheus.yml 
./prometheus --config.file=prometheus.yml
```

## alertmanager-安装

### Ubuntu安装

```bash
# 访问：https://prometheus.io/download/
# 下载组件prometheus/alertmanager：alertmanager-0.23.0.linux-amd64.tar.gz
# 解压到目标文件夹：alertmanager
# 到目标文件夹中
# 启动：
./alertmanager
# 访问：http://localhost:9093/

# 使用配置文件启动:配置alertmanager.yml
./alertmanager --config.file=alertmanager.yml
```