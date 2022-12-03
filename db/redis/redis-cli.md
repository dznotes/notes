# redis-cli

## redis-cli-参数示例

## redis-cli-使用示例

- 简单示例

```bash
# 本地+没有密码
redis-cli

# 本地+有密码
redis-cli -a password

# 远程+有密码
redis-cli -h 192.168.1.x -p 6379 -a password
```

### 查看大key

```bash
# 可能依赖于dbsize
# 查看大key
redis-cli --bigkeys

# 查看大key：每扫描100个key休息0.1秒
redis-cli --bigkeys -i 0.1
```
### 清除数据

```bash
# 清除数据
redis-cli keys "xx*" | xargs redis-cli del
```

### 批量删除待空格的key

```bash
redis-cli  keys "xx_*" > a.txt

awk '$0="redis-cli xx del \""$0"\""' a.txt > cmd.txt

chmod a+x cmd.txt
./cmd.txt
```