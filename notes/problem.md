https://redis.io/docs/manual/patterns/distributed-locks/
redis分布式锁问题:
set key value nx px 1000
需要设置过期时间,不设置过期时间会出现deadlock现象(在释放锁(del key)之前client挂掉了)
需要设置unique value安全删除锁,clientA在运行过程中长时间block,锁的过期时间到了被redis回收,clientB获取到了锁,clientA会在不知情的情况下删除锁
安全删除锁you两步(check value == ***;delete key)通过lua script完成
如果单实例的redis挂了,锁将会丢失
如果redis采用m-s的方式做redis的高可用,也无法保证锁的有效性(redis的replica是异步的,如果锁在创建之后在还未replica到slave之前,master down,slave prompt,此时slave并没有锁的key)

阿里云的redis服务对于多key的cmd可能不太支持,需要阅读使用手册

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