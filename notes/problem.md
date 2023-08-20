

https://www.youtube.com/watch?v=f7VsHLk_Z8c
TPS/QPS
throughput/lantency

lantency(network processing) how fast
throughput how much

Parameters of system Throughput

- QPS/TPS (Number of request/query/transactions per second): The number of requests/transactions per second
- Concurrency: The number of requests/transactions processed by the system at the same time
- Response Time: The time to complete a request (Generally take the average response time)
After understanding the meaning of the above 3 elements, the relationship between them can be deduced:

**QPS/TPS = concurrency/average response time**
max throughput = threads/latency

UV
VV
PV


当然，现实中 50000 QPS 的系统几乎必然拥有负载均衡器，即便每个接口只返回 20KB，那网络带宽也已经达到了 976MB/S，即 7.8Gbit，单机带宽都快干到 10G 了，肯定是不会只用单台服务器硬抗的，即使单机性能能达到，那单机也无法保证这么大规模系统的稳定性。这个时候我们就需要负载均衡器的介入，接下来两篇文章我们会详细讨论。


排行榜的具体实现  scan redis会有什么问题
golang实现 default param