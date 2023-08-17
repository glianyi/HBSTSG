CPU, processors, core, threads

#### 一些概念
question:
- IO密集型是用多线程还是多进程效率高些，如果是单核的CPU机器上呢
- 如果是CPU密集型下使用多线程还是多进程效率高些
阅读材料
- https://www.golinuxcloud.com/goroutine-vs-threads-golang/
- io密集用多线程或者io多路复用, I/O密集型:当线程等待时间所占比例越高,需要越多线程,启用其他线程继续使用CPU,以此提高CPU的利用率;
- CPU密集型:当线程CPU时间所占比例越高,需要越少的线程.任务越多,花在进程、线程切换的时间就越多,通常线程数和CPU核数一致即可,这一类型在开发中主要出现在一些计算业务频繁的逻辑中

#### memory leak
阅读材料
- https://walterwu-22843.medium.com/golang-memory-leak-while-handling-huge-amount-of-http-request-35cc970cb75e

- 分析内存的工具
- global var eg. var globalSlice = make([]int,0),slice的占用空间会一直变大
- 持续增长的goroutine数量(太多hanging的goroutine或者goroutine leak)(Usually, a hanging Go routine is not as simple as the sleeper example above. It can be an http client that keeps connection alive. It can also be a dead loop. Long polling or web socket client both keep connection open forever. If you are going to use a never-ending Go routine, make sure there is only one such connection.)
- open stream(打开文件之后未及时关闭文件)
- 本地复现memory leak

#### tips
- Always close the request.body

#### 内存对齐
阅读材料
question:
- 操作系统为什么要做一些内存对齐(内存对齐在某些情况下可以减少读取内存的次数以及一些运算，性能更高,内存对齐提升性能的同时,也需要付出相应的代价.由于变量与变量之间增加了填充,并没有存储真实有效的数据,所以占用的内存会更大.这也是一个典型的空间换时间的应用场景)
- cpu读取内存数据是按"箱读取",[0,1,2,3][4,5,6,7]当读取起始地址为3的数据时cpu需要从0开始读取完整的[0,1,2,3,4],如果数据跨箱子了就需要读取多次,读取跨箱子的数据还会增加cpu指令,原子性得不到保证

#### 并发与并行
- 并行是特殊的并发,同一个cpu core上可以实现并发,但是无法实现并行

#### go语言的并发机制与模型
questions:
- 如果goroutine一直占用资源。GPM模型是如何处理的(如果G占用cpu操作10ms会被调度器抢占)
- go协程调度和系统调度有什么区别
- go语言的并发,并发模型的理解
- GPM模型，P存在的意义是什么(P代表一组goroutine同M绑定,同源的G可能需要再同一个M上执行)
- 一个G阻塞后，会发生什么；等待结束之后如何唤醒（G阻塞之后,与之对应的M也会阻塞,M会让出P,阻塞结束之后G会回到全局队列等待新的M
- go调度原理


阅读资料
- https://www.syntio.net/en/labs-musings/gos-concurrency-model/
- https://learnku.com/articles/41728
- GMP(G:goroutine M:os thread P:Processor)
- GM(N:M) MP(1:1)P的大小是runtime启动之后就确定的(GOMAXPROCS),当P的数量多于M的数量的时候会创建M,当P中的G阻塞去寻找空闲的M,如果找不到就会创建M
- N goroutines are scheduled onto M operating system threads,N is usually bigger than M, and M is usually set to the number of CPU cores (the default value in Go), because it allows for maximum paralellization. One very important thing to realize, which is the whole basis of this concept, is that a sleeping (blocked) goroutine doesn’t waste CPU time.

调度器的设计策略
- 复用M(M空闲(对应的P中无goroutine)之后不立即销毁,从其他P中获取;当M因goroutine阻塞的时候解绑P,这样P可以被其他M绑定)
- 并行:GOMAXPROCS设置P的数量,最多有GOMAXPROCS个线程分布在多个CPU上同时运行.GOMAXPROCS也限制了并发的程度;
- 抢占:一个 goroutine最多占用CPU 10ms,防止其他goroutine被饿死;
- 全局G队列:在新的调度器中依然有全局G队列,但功能已经被弱化了,当M执行work stealing从其他P偷不到G时,它可以从全局G队列获取G;

GMP:
- thread与goroutine的绑定机制
-- 如果所有goroutine都绑定在一个thread上,其中一个goroutine阻塞导致了thread的阻塞,那该thread上的其他thread也都会阻塞,同时也无法利用多core的能力
-- 如果一个goroutine对应一个thread,可以利用上多核,gorputine的调度就退化成了thread的调度
-- goroutine的调度同thread的调度是由区别的,thread是os抢占式调度(cpu始终周期驱动),goroutine是runtime scheduler进行协作调度(一个goroutine 让渡出thread之后才能被下一个goroutine使用,go scheduler也会有抢占策略,一个goroutine最多占用CPU10ms,防止其他 goroutine 被饿死)


goroutine:
- 属性stack(2kb) thread(2M)
- 生命周期

#### process thread goroutine
- 都解决了阻塞带来的cpu浪费问题,process thread切换的代价大,如果多了cpu用在切换上的占比就高了,同时内存消耗也大

### view
context的实现、这个数据结构是并发安全的吗
- http request,exec command
- https://p.agnihotry.com/post/understanding_the_context_package_in_golang/
- context是基于goroutine实现的,是并发安全的
- TODO() 不知道传什么样的ctx的时候就用这个 Background() top level

waitGroup是如何做到所有的协程都done后，wait函数才会去执行、内部计数器对数据的加减是并发安全的吗
- https://ferencfbin.medium.com/golang-source-1-waitgroup-under-the-hood-281ea6101f54
- add() add() wait() ok add() wait() add() not ok
- waitGroup里边有一个state变量,只有当state变量为zero的是后 wait才会执行,state变量的操作是原子操作(atomic)并发安全的

channel:
- channel优势、为什么是线程安全的
- channel无缓冲和有缓冲有什么区别
- channel实现(https://stackoverflow.com/questions/19621149/how-are-go-channels-implemented https://docs.google.com/document/d/1yIAYmbvL3JxOKOjuCyon7JhW4cSv1wy5hC0ApeGMV9s/pub)
- goroutine泄露 原理以及什么场景会发生goroutine泄露
- 对channel的操作什么情况下会出现死锁
- channel为什么是线程安全的(channel只有在多个goroutine中才能正常使用,channel内部数据结构是有mutex控制并发安全的)
- https://go.dev/ref/spec#Close
- https://stackoverflow.com/questions/34897843/why-does-go-panic-on-writing-to-a-closed-channel
- https://groups.google.com/g/golang-nuts/c/_Q6FCjWIr18/m/Wb_NlyusRL8J
- Sending to or closing a closed channel causes a run-time panic
- Closing the nil channel also causes a run-time panic
- close(channel)类似于free(pointer),可能会回收一些资源,如果多次free\close会导致内部数据结构破坏;close(ch)实际上是向channel发送一个特殊值(如果是nil channel就会panic)
- channel两端丢失任何一端都会死锁 or 出现在同一个goroutine
- channel导致的goroutine,channel两端的其中一端的goroutine结束了,另一端会一直block在channel上
- channel的生产者不发送channel数据了,消费者(for range channel)会一直阻塞,当生产者channel调用close(ch) 消费者(for range)会受到一个信号解除阻塞
- 创建goroutine的时候需要关注goroutine的生命周期,当goroutine结束的时候需要做一些资源回收或者通知(close(ch))操作
- channel两边没有生产消费者的时候 gc会回收channel;只有当channel其中一方需要通知另一方结束这个场景才会主动关闭channel
- 不要再reciever这一端close(ch),可能会导致sender send to closed channel lead to panic
- 如果一个channel有多个sender,尽量不要close(ch),可能会导致sender sender to closed channel or close closed channel
- https://go101.org/article/channel-closing.html

channel的sender,reciever的几种模式
- 1s1r: sender close channel
- 1snr: sender close channel
- nsnr: sender之前通过一个channel协调控制close channel



#### 锁
- 读写锁和互斥锁的区别(读写锁:写互斥;互斥锁:读写都互斥)

#### map
sync.Map如何实现的，更新一个key的过程

##### gin
gin的中间件的实现
gin的路由匹配规则是怎么实现的、有针对于前缀树做什么优化没
map 对key遍历是如何做的

从几千万表里随机查找符合范围的数据
如何保证MySQL 和redis 数据一致性
机器4g，突然起了一个1g的线程，从性能角度分析一下
select和epoll的区别
redis分布式锁
redis缓存穿透,缓存雪崩
时序数据库如何实现
raft选举,高可用
kafka的offset怎么管理
kafka多个consumer group 消费同一个 partition 有问题吗？为什么同一个 consumer group 里不能消费同一个partition
数据库中的乐观锁悲观锁
程序中的乐观锁和悲观锁
消息队列消费端的推和拉有什么区别
长链接转短链接设计

作业帮(二面)

#### 长连接保持问题
前端gateway集群维护客户端长连接,gateway分发请求到后端集群处理逻辑后,如何保证后端处理完毕回包回到正确的gateway机器
承接第一个问题,如果机器挂掉,如何最小损失情况下恢复,k8s集群
https://medium.com/lumen-engineering-blog/how-to-implement-a-distributed-and-auto-scalable-websocket-server-architecture-on-kubernetes-4cc32e1dfa45

有哪些情况会发生逃逸
栈空间为什么会不足，为什么会逃逸



写一个接口，调用3个服务，并发调用接口，并且API超时报错，结束本次请求
写一个死锁SQL，并分析
给定一个数组，并发交替打印奇数，偶数，不能使用锁，原子操作

#### view
- 对map并发读写可能会panic
- map为什么无序(map的遍历是内部数据结构bucket的遍历,bucket的遍历是随机的并不一定从0号bucket开始,bucket内部的kv遍历是有序的但是扩容之后key可能到其他的- bucket了)
- sync.Map为什么无序
- 传入一个slice，函数内部对他进行修改，slice改变与否(会 slice map都会)(基础类型、array\struct不会)
实现一个有序map、支持add、delete、迭代(how)

java里的interface和go里的interface有什么区别
gpm调度
起100个goroutine，其中50个sleep 10s，具体的调度过程

goroutine数据结构里面都存了啥字段
goroutine有哪些状态，如何扭转
聊聊Go  GC
go最近几个版本有什么新的变化、项目升级了吗、为什么选择升/不升

defer在最近几个版本是如何做的性能提升
并发若干个goroutine、其中一个panic怎么办
了解go的内存对齐么，谈一谈

nsq原理、实现、以及如何保证消息不丢



在浏览器输入一个网址，数据是如何到服务端的。经历了哪些过程
HTTPS通信步骤
HTTPS和http区别。是如何保证安全的
浏览器崩溃了、是客户端还是服务端断开连接的
client是如何实现长连接的

string和byte是啥关系
byte和rune区别
互斥锁如何实现公平

map实现原理、如何实现读、写
map扩容
sync.map 如何保证线程安全


Go中的http包的实现原理
Golang 里的逃逸分析是什么
Go中对nil的Slice和空Slice的处理
Goroutine 如何调度?
go内存对齐
map 如何实现、原理

逃逸分析
hello Goroutine的执行过程
Go语言的栈空间管理是怎么样的
Goroutine和线程的区别
sync.map是如何实现的
channel的使用场景
context 超时有用么

http 2.0 （以及3.0）和1.0 区别
描述dns解析