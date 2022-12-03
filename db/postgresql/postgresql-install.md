# postgresql-安装教程

## postgresql-安装

### Ubuntu安装

```bash
sudo apt-get install postgresql postgresql-contrib

pg_ctlcluster 12 main start
sudo systemctl start postgresql@12-main

# 查看pg状态
service postgresql status

# 切换用户
sudo su postgres

# 使用shell
psql

\l 	# 查看现有的所有表
\du # 查看postgresql 用户
```