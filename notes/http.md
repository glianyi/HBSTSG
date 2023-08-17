```http
GET / HTTP/1.0
Content-Type: application/json

{
    "code": 200
}
```
请求行: verb uri version
http请求可以没有body只有HEAD(HEAD无限制,取决于server的具体实现)

响应
状态行: version code reason(http/1.0 200 ok)

http的request,response都包含了HEADER,header定义了整个http数据包,content-length指定了body的长度, \r\n\r\n分割了HEADER与body 如果HEADER多了一个\r\n,会当做body的一部分处理

HEADER与query参数(key=value)的职责划分
如果只有key就表示没有value
http是可靠协议,但是不能保证数据百分之百被送达,如果要保证数据被送达需要使用类似kafka这样的协议
http是无连接的,但是为了保证性能1.1增加了keepalive选项保持连接

语义上GET HEAD PUT DELETE是幂等,post不幂等

uri: scheme://user:password@host:port/path?query#fragment

uri中只允许出现ASICII编码,如果出现非ASICII编码则直接escape(转义成"%十六进制值")
