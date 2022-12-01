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