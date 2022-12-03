# rocketmq-安装教程

## rocketmq-安装

### window安装

```base
# 1、下载：
https://mirrors.tuna.tsinghua.edu.cn/apache/rocketmq/4.8.0/rocketmq-all-4.8.0-bin-release.zip
# 2、解压：
解压到C:\rocketmq


## 环境变量配置
ROCKETMQ_HOME="C:\rocketmq"
NAMESRV_ADDR="localhost:9876"
```

### docker安装

```bash
docker search rocketmq
docker pull rocketmqinc/rocketmq


# 启动Nameserver
# 方式2
// https://blog.csdn.net/weixin_38836909/article/details/103985478
// -d: 后台运行容器，并返回容器ID；
// -p: 指定端口映射，格式为：主机(宿主)端口:容器端口。
// --name="nginx-lb": 为容器指定一个名称；
$ docker run -d -p 9876:9876 \
--name rmqserver  \
foxiswho/rocketmq:server-4.5.1

# 启动Broker
# 对应上面方式2
// Broker容器中默认的配置文件的路径为：/etc/rocketmq/broker.conf
$ docker run -d -p 10911:10911 -p 10909:10909 \
--name rmqbroker \
--link rmqserver:namesrv\
-e "NAMESRV_ADDR=namesrv:9876"\
-e "JAVA_OPTS=-Duser.home=/opt"\
-e "JAVA_OPT_EXT=-server -Xms128m -Xmx128m"\
foxiswho/rocketmq:broker-4.5.1



# 安装控制台
docker search rocketmq-console
docker pull styletang/rocketmq-console-ng　

# 方式2
// 访问 http://localhost:8082/#/
$ docker run -d \
--name rmqconsole \
-p 8082:8080 \
--link rmqserver:namesrv\
-e "JAVA_OPTS=-Drocketmq.namesrv.addr=namesrv:9876\
-Dcom.rocketmq.sendMessageWithVIPChannel=false"\
-t styletang/rocketmq-console-ng
```