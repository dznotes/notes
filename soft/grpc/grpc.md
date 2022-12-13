# grpc

## grpc-文档参考
### 网站
- 官网：https://www.grpc.io/
- awesome-grpc：https://github.com/grpc-ecosystem/awesome-grpc
- 官方文档-go：https://grpc.io/docs/languages/go/
- github：https://github.com/grpc/grpc-go

### 中间件

- https://github.com/grpc-ecosystem/go-grpc-middleware
- 限流器：https://github.com/grpc-ecosystem/go-grpc-middleware/blob/master/ratelimit/ratelimit.go
- 文章：https://eddycjy.com/posts/go/grpc/2018-10-10-interceptor/

### 书籍/教程

- https://www.selinux.tech/golang/grpc
- https://www.cnblogs.com/FireworksEasyCool/category/1693727.html
- https://github.com/Bingjian-Zhu/go-grpc-example
- gRPC与云原生应用开发
    - https://grpc-up-and-running.github.io/
    - 示例代码：https://github.com/grpc-up-and-running/samples

## grpc-组件之间的关系

## grpc-错误码

## grpc-基础概念

### 基础通信模式（4种）

| No.  | 模式        | 解释                         |
| :--- | :---------- | :--------------------------- |
| 01   | 一元RPC     | 简单RPC模式。                |
| 02   | 服务端流RPC | 客户端1条消息+服务端流序列。 |
| 03   | 客户端流RPC | 客户端流序列+服务端1条消息。 |
| 04   | 双向流RPC   | 客户端流序列+服务端流序列。  |

### 多路复用

- 单个客户端 存根 使用gRPC客户端连接，还允许在**同一个gRPC服务器端运行多个gRPC服务**。
- 也允许**多个客户端 存根 使用同一个gRPC客户端连接**。
- 这个功能叫做 **多路复用**。

### 元数据

- metadata：在客户端和服务端之间传递信息。

### 负载均衡

- **负载均衡器代理：客户端与服务器之间存在 1个负载均衡器，客户端请求负载均衡器；服务器将负载情况报告给负载均衡器。**
- **客户端负载均衡：客户端层面做负载均衡。**
- **轮询调度算法（round-robin algorithm）: pick_first（第一个成功地址），round_robin（轮询）。**

## grpc-常用拦截器

- 功能：认证、日志、消息、校验、重试、监控

| 分类 | 拦截器           | 介绍                                   |
| ---- | ---------------- | -------------------------------------- |
| 认证 | grpc_auth        | 可自定义的认证中间件                   |
| 日志 | grpc_ctxtags     | 添加Tag map到上下文的库                |
|      | grpc_zap         | 将zap日志库集成到grpc handler中        |
|      | grpc_logrus      | 将logrus日志库集成到grpc handler中     |
| 监控 | grpc_prometheus  | prometheus 客户端和服务端的监控中间件  |
|      | grpc_opentracing |                                        |
| 重试 | grpc_retry       | 客户端：gRPC响应码重试机制             |
| 验证 | grpc_validator   | 服务端：根据.proto选项生成入站消息校验 |
|      | grpc_recovery    | 服务端：将panic转换为gRPC错误          |
|      | ratelimit        | 服务端：限制gRPC的速度                 |

## grpc-常用错误码

| 错误码              | 数字 | 描述                                       |
| ------------------- | ---- | ------------------------------------------ |
| ok                  | 0    | 成功状态                                   |
| cancelled           | 1    | 操作已被（调用者）取消                     |
| unknown             | 2    | 未知错误                                   |
| invalid_argument    | 3    | 客户端指定了非法参数                       |
| deadline_exceeded   | 4    | 在操作完成前，就已超过了截止时间           |
| not_found           | 5    | 某些请求实体没有找到                       |
| already_exists      | 6    | 客户端试图创建的实体已经存在               |
| permission_denied   | 7    | 调用者没有权限执行特定的操作               |
| resource_exhausted  | 8    | 某些资源已经被耗尽                         |
| failed_precondition | 9    | 操作被拒绝，系统没有处于执行操作所需的状态 |
| aborted             | 10   | 操作被中止                                 |
| out_of_range        | 11   | 尝试进行的操作超出了合法的范围             |
| unimplemented       | 12   | 在该服务中，未实现或不支持本操作           |
| internal            | 13   | 内部错误                                   |
| unavailable         | 14   | 该服务当前不可用                           |
| data_loss           | 15   | 不可恢复的数据丢失或损坏                   |
| unauthenticated     | 16   | 客户端没有进行操作的合法认证凭证           |

## grpc-网关

- github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
- github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

```protobuf
import "google/api/annotations.proto";

rpc add() returns () {
	option (google.api.http) = {
  	post: "/v1/product"
    body: "*"
  };
}
```

