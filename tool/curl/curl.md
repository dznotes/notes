# curl
## 使用示例
```bash
curl -o /dev/null -s -w  "\n"HTTP返回码：%{http_code}"\n"对端IP地址：%{remote_ip}"\n"应用建连时间：%{time_appconnect}"\n"TCP连接时间：%{time_connect}"\n"DNS解析时间：%{time_namelookup}"\n"准备传输时间：%{time_pretransfer}"\n"开始传输时间：%{time_starttransfer}"\n""\n"总时间：%{time_total}"\n"  \
www.baidu.com
```