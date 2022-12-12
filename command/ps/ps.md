# ps
- Process Status

## ps-命令参数

| 命令  | 含义         |
| :---- | :----------- |
| ps -A | 显示所有进程 |

## ps-输出含义

| 字段    | 含义                               |
| :------ | :--------------------------------- |
| USER    | 用户名称                           |
| UID     | 用户ID / User ID                   |
| PID     | 进程ID / Process ID                |
| PPID    | 父进程的进程ID / Parent Process ID |
| %CPU    | 进程的CPU占用率                    |
| %MEM    | 进程的内存占用率                   |
| TTY     | 与进程关联的终端 tty               |
| TIME    | 进程使用的总CPU时间                |
| COMMAND | 正在执行的命令行命令               |



## ps-常用命令

| 命令                | 含义                              |
| :------------------ | :-------------------------------- |
| ps                  |                                   |
| ps -ef              | 显示出linux机器所有详细的进程信息 |
| ps -ef  &vert; grep xx | 查找特定进程                      |
| ps -u xxx           | 查看指定用户进程：ps -u root      |
