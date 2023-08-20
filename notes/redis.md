配置redis代码阅读环境(vscode)
vscode打开代码目录,找到main(server.c)文件,点击运行,生成.vscode目录(c_cpp_properties.json,launch.json,settings.json,task.json)
c_cpp_properties
```json
{
    "configurations": [
        {
            "name": "Linux",
            "includePath": [
                "${workspaceFolder}/src/**",
                "${workspaceFolder}/deps/lua/**",
                "${workspaceFolder}/deps/hiredis/**",
                "${workspaceFolder}/deps/hdr_histogram/**"
            ],
            "intelliSenseMode": "linux-gcc-x64",
            "compilerPath": "/usr/bin/gcc",
            "cStandard": "c11",
            "cppStandard": "c++14"
        }
    ],
    "version": 4
}
```
launch.json
```json
{
    // 使用 IntelliSense 了解相关属性。 
    // 悬停以查看现有属性的描述。
    // 欲了解更多信息，请访问: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "type": "cppdbg",
            "request": "launch",
            "name": "redis launch",
            "program": "${workspaceFolder}/src/redis-server",
            "args": ["redis.conf"],
            "cwd": "${workspaceFolder}",
            "MIMode": "gdb"
        }
    ]
}
```

settings.json
```json
{
    "files.associations": {
        "monotonic.h": "c"
    }
}
```

tasks.json
```json
{
    "tasks": [
        {
            "type": "shell",
            "label": "redis-build",
            "command": "/usr/bin/make",
            "args": [],
            "options": {
                "cwd": "${workspaceFolder}"
            },
            "problemMatcher": [
                "$gcc"
            ],
            "group": "build",
            "detail": "调试器生成的任务。"
        }
    ],
    "version": "2.0.0"
}
```

https://redis.io/docs/manual/patterns/distributed-locks/
redis分布式锁问题:
set key value nx px 1000
需要设置过期时间,不设置过期时间会出现deadlock现象(在释放锁(del key)之前client挂掉了)
需要设置unique value安全删除锁,clientA在运行过程中长时间block,锁的过期时间到了被redis回收,clientB获取到了锁,clientA会在不知情的情况下删除锁
安全删除锁you两步(check value == ***;delete key)通过lua script完成
如果单实例的redis挂了,锁将会丢失
如果redis采用m-s的方式做redis的高可用,也无法保证锁的有效性(redis的replica是异步的,如果锁在创建之后在还未replica到slave之前,master down,slave prompt,此时slave并没有锁的key)

阿里云的redis服务对于多key的cmd可能不太支持,需要阅读使用手册

单实例redis支持的最大连接数 maxclient

列表中的数据是否压缩
- 元素长度<48 或者压缩前后差别不大

列表应用 lrange lists 0 -1
栈(LPUSH LPOP)
队列(LPUSH RPOP)
阻塞队列(LPUSH BRPOP),阻塞的含义(一个连接过来执行pop,队列是空的就会阻塞,同一个连接后面的命令也会被阻塞)
延时队列
异步消息队列(消除svc svc之间的长连接)
获取固定窗口记录(限流 lpush ltrim维护一个窗口)

hash应用(缓存用户原信息) hashmap
hset id:user name mark age 10 gendar male

set应用 intset  唯一性 无序性 (好友共同关注[交集] 推荐关注[补集])
zset 有序集合  时序数据 排行榜


redis网络层
redis异步连接 单reactor(epoll) 一个线程同时处理 指令和网络io mysql是如何处理连接的
对于连接的处理redis是并发 mysql也是并发(受限于主机的核心数)
对于命令的处理(同一个连接)redis是串行 mysql是并发
redis事务(ACID)(multi exec discard watch)
redis pipeline(是客户端技术,客户端使用什么样的技术(阻塞 非阻塞)同redis server通信)
redis发布订阅

一个连接对应一个读缓冲和写缓冲 客户端只是将数据写到了连接对应的缓冲区,由内核协议栈实现发送到对端

redis集群
https://medium.com/opstree-technology/redis-cluster-architecture-replication-sharding-and-failover-86871e783ac0

https://medium.com/the-quarter-espresso/introduction-of-neural-redis-part-1-fa13c1faeef1

io密集与计算密集



source code
main -> initListeners -> createSocketAcceptHandler -> aeCreateFileEvent

initServer -> aeCreateEventLoop,aeCreateTimeEvent,aeCreateFileEvent

get foo
*2\r\n$3\r\nget\r\n$3\r\nfoo\r\n

客户端发起连接会发送("*2\r\n$7\r\nCOMMAND\r\n$4\r\nDOCS\r\n\301\017\302\366\377\177")

keys *
21 "*2\r\n$4\r\nkeys\r\n$1\r\n*\r\n"

\r\nkey_spec_index\r\n:1\r\n$5\r\nflags\r\n*1\r\n+multiple\r\n*10\r\n$4\r\nname\r\n$6\r\nweight\r\n$4\r\ntype\r\n$7\r\ninteger\r\n$12\r\ndisplay_text\r\n$6\r\nweight\r\n$5\r\ntoken\r\n$7\r\nWEIGHTS\r\n$5\r\nflags\r\n*2\r\n+optional\r\n"

https://stackoverflow.com/questions/8911883/debugging-of-a-c-program-redis-server

#### how redis event loop works
redis事件循环
https://draveness.me/redis-eventloop/
https://redis.io/docs/reference/internals/internals-rediseventlib/
https://app.codecrafters.io/walkthroughs/redis-event-loop