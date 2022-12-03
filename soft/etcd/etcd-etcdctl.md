# etcdctl使用

## etcdctl-文档参考


## etcdctl-使用

- 注意：etcdctl_api 2和3的api数据不互通


### API=2

```bash 
export ETCDCTL_API=2

# 简单的增（更）-查-删
etcdctl put sam test
etcdctl get sam
etcdctl del sam

# 前缀查询
etcdctl put /class/A/A1 01 
etcdctl put /class/A/A2 02
etcdctl put /class/B/B1 04
etcdctl put /class/B/B2 04

etcdctl get /class --prefix
etcdctl get /class/A --prefix

# watch机制
etcdctl watch /class --prefix
etcdctl put /class/A/A1 100
etcdctl put /class/C/C1 200
```

### API=3

```bash
export ETCDCTL_API=2
# API version = 2
# 增删改查
etcdctl set sam test
etcdctl get sam
etcdctl update sam 200
etcdctl get sam

# 成员查看
etcdctl member list
```