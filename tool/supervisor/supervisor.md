# supervisor

## supervisor-文档参考

- supervisord 是 server 端。 
- supervisorctl 是 client 端。

## supervisor-安装

### mac安装

```bash

```

### ubuntu安装

```bash
sudo apt-get install supervisor
```

## supervisorctl-常用命令



| 命令                      | 含义                                         |
| :------------------------ | :------------------------------------------- |
| supervisorctl start xxx   | 启动进程                                     |
| supervisorctl restart xxx | 重启进程                                     |
| supervisorctl stop group  | 停止所有属于名为group的分组进程              |
| supervisorctl stop all    | 停止全部进程                                 |
| supervisorctl reload      | 载入最新配置的文件                           |
| supervisorctl update      | 根据最新的配置文件，启动新配置或有改动的进程 |



## supervisor-conf配置文件

- 创建文件，写入配置

```bash
cd /etc/supervisor/conf.d/
sudo touch test.conf

[program:test]
command=sh /usr/local/bin/test.sh                  ;被监控的进程路径
numprocs=1                    ; 启动一个进程
directory=/usr/local/bin/     ;执行前切换路径
autostart=true                ; 随着supervisord的启动而启动
autorestart=true              ; 自动重启
startretries=10               ; 启动失败时的最多重试次数
exitcodes=0                   ; 正常退出代码
stopsignal=KILL               ; 用来杀死进程的信号
stopwaitsecs=10               ; 发送SIGKILL前的等待时间
redirect_stderr=true          ; 重定向stderr到stdout
stdout_logfile=logfile        ; 指定日志文件
```



- 配置项解释

| 命令        | 解释                                | 示例             |
| ----------- | ----------------------------------- | ---------------- |
| autorestart | 程序异常退出后自动重启              | autorestart=True |
| autostart   | 在 supervisord 启动的时候也自动启动 | autostart=True   |
|             |                                     |                  |

