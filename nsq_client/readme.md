#### nsq 相关文档

1. https://zhuanlan.zhihu.com/p/37081073

- 重点摘要

channel消息传递的通道:

1. 当生产者每次发布消息的时候,消息会采用多播的方式被拷贝到各个 channel 中, channel 起到队列的作用。
2. channel 与 consumer(消费者) 相关，是消费者之间的负载均衡,消费者通过这个特殊的channel读取消息。
3. 在 consumer 想单独获取某个 topic 的消息时，可以 subscribe(订阅)一个自己单独命名的 nsqd中还不存在的 channel, nsqd会为这个 consumer创建其命名的 channel
4. Channel 会将消息进行排列，如果没有 consumer读取消息，消息首先会在内存中排队，当量太大时就会被保存到磁盘中。可以在配置中配置具体参数。
5. 一个 channel 一般会有多个 consumer 连接。假设所有已连接的 consumer 处于准备接收消息的状态，每个消息将被传递到一个随机的 consumer。
6. Go语言中的channel是表达队列的一种自然方式，因此一个NSQ的topic/channel，其核心就是一个存放消息指针的Go-channel缓冲区。缓冲区的大小由 --mem-queue-size 配置参数确定。

consumer消息的消费者:

1. consumer 通过 TCPsubscribe 自己需要的 channel
2. topic 和 channel 都没有预先配置。 topic 由第一次发布消息到命名 topic 的 producer 创建 或 第一次通过 subscribe 订阅一个命名 topic 的 consumer 来创建。 channel
   被 consumer 第一次 subscribe 订阅到指定的 channel 创建。
3. 多个 consumersubscribe一个 channel，假设所有已连接的客户端处于准备接收消息的状态，每个消息将被传递到一个 随机 的 consumer。
4. NSQ 支持延时消息， consumer 在配置的延时时间后才能接受相关消息。
5. Channel在 consumer 退出后并不会删除，这点需要特别注意。

### nsqadmin参数

```text
Message Queues:

Depth: ：内存和硬盘转存的消息数.
In-Flight: 当前未完成的消息数：包括发送但未返回FIN/重新入队列REQ/超时TIMEOUT 三种消息数之和.
Deferred: 当前重入队列但还没重新发送消息的数量，这个字段作用是？
Statistics:

Requeued: Total count of messages that have been added back to the queue due to time outs or explicit requeues. 重入队列数（或者因为超时重入队列，或者明确要求重入队列）
Timed Out: Total count of messages that were requeued after not receiving a response from the client before the configured timeout. 已重入队列但按配置的超时时间内还收到响应的消息数
Messages: Total count of new messages recieved since node startup. 节点起来后的所有新消息总数，真正的消息次数
Rate: The per-second rate of new messages (available when graphite integration is enabled). 每秒的消息数
```

#### nsqd命令行参数

```text
-auth-http-address=: <addr>:<port> 查询授权服务器 (可能会给多次)
-broadcast-address="": 通过 lookupd  注册的地址（默认名是 OS）
-config="": 配置文件路径
-data-path="": 缓存消息的磁盘路径
-deflate=true: 运行协商压缩特性（客户端压缩）
-e2e-processing-latency-percentile=: 消息处理时间的百分比（通过逗号可以多次指定，默认为 none）
-e2e-processing-latency-window-time=10m0s: 计算这段时间里，点对点时间延迟（例如，60s 仅计算过去 60 秒）
-http-address="0.0.0.0:4151": 为 HTTP 客户端监听 <addr>:<port>
-https-address="": 为 HTTPS 客户端 监听 <addr>:<port>
-lookupd-tcp-address=: 解析 TCP 地址名字 (可能会给多次)
-max-body-size=5123840: 单个命令体的最大尺寸
-max-bytes-per-file=104857600: 每个磁盘队列文件的字节数
-max-deflate-level=6: 最大的压缩比率等级（> values == > nsqd CPU usage)
-max-heartbeat-interval=1m0s: 在客户端心跳间，最大的客户端配置时间间隔
-max-message-size=1024768: (弃用 --max-msg-size) 单个消息体的最大字节数
-max-msg-size=1024768: 单个消息体的最大字节数
-max-msg-timeout=15m0s: 消息超时的最大时间间隔
-max-output-buffer-size=65536: 最大客户端输出缓存可配置大小(字节）
-max-output-buffer-timeout=1s: 在 flushing 到客户端前，最长的配置时间间隔。
-max-rdy-count=2500: 客户端最大的 RDY 数量
-max-req-timeout=1h0m0s: 消息重新排队的超时时间
-mem-queue-size=10000: 内存里的消息数(per topic/channel)
-msg-timeout="60s": 自动重新队列消息前需要等待的时间
-snappy=true: 打开快速选项 (客户端压缩)
-statsd-address="": 统计进程的 UDP <addr>:<port>
-statsd-interval="60s": 从推送到统计的时间间隔
-statsd-mem-stats=true: 切换发送内存和 GC 统计数据
-statsd-prefix="nsq.%s": 发送给统计keys 的前缀(%s for host replacement)
-sync-every=2500: 磁盘队列 fsync 的消息数
-sync-timeout=2s: 每个磁盘队列 fsync 平均耗时
-tcp-address="0.0.0.0:4150": TCP 客户端 监听的 <addr>:<port>
-tls-cert="": 证书文件路径
-tls-client-auth-policy="": 客户端证书授权策略 ('require' or 'require-verify')
-tls-key="": 私钥路径文件
-tls-required=false: 客户端连接需求 TLS
-tls-root-ca-file="": 私钥证书授权 PEM 路径
-verbose=false: 打开日志
-version=false: 打印版本
-worker-id=0: 进程的唯一码(默认是主机名的哈希值)

```

#### nsqlookupd 命令行选项

```text

     -http-address="0.0.0.0:4161": <addr>:<port> 监听 HTTP 客户端
     -inactive-producer-timeout=5m0s: 从上次 ping 之后，生产者驻留在活跃列表中的时长
     -tcp-address="0.0.0.0:4160": TCP 客户端监听的 <addr>:<port>
     -broadcast-address: 这个 lookupd 节点的外部地址, (默认是 OS 主机名)
     -tombstone-lifetime=45s: 生产者保持 tombstoned  的时长
     -verbose=false: 允许输出日志
     -version=false: 打印版本信息
```

#### nsqadmin命令行选项

```text
-graphite-url="": URL to graphite HTTP 地址
-http-address="0.0.0.0:4171": <addr>:<port> HTTP clients 监听的地址和端口
-lookupd-http-address=[]: lookupd HTTP 地址 (可能会提供多次)
-notification-http-endpoint="": HTTP 端点 (完全限定) ，管理动作将会发送到
-nsqd-http-address=[]: nsqd HTTP 地址 (可能会提供多次)
-proxy-graphite=false: Proxy HTTP requests to graphite
-template-dir="": 临时目录路径
-use-statsd-prefixes=true: expect statsd prefixed keys in graphite (ie: 'stats_counts.')
-version=false: 打印版本信息
```

