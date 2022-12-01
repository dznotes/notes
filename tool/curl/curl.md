# curl
## 使用示例
```bash
curl -o /dev/null -s -w  "\n"HTTP返回码：%{http_code}"\n"对端IP地址：%{remote_ip}"\n"应用建连时间：%{time_appconnect}"\n"TCP连接时间：%{time_connect}"\n"DNS解析时间：%{time_namelookup}"\n"准备传输时间：%{time_pretransfer}"\n"开始传输时间：%{time_starttransfer}"\n""\n"总时间：%{time_total}"\n"  \
www.baidu.com

# 输出 示例
HTTP返回码：200
对端IP地址：110.242.68.4
应用建连时间：0.000000
TCP连接时间：0.021064
DNS解析时间：0.008310
准备传输时间：0.021111
开始传输时间：0.034913

总时间：0.035058
```