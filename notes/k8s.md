authentication(https://gcore.com/learning/kubernetes-authentication/)
Kubernetes assumes that ‘users’ are managed outside of Kubernetes:
     a. In production environments it can be LDAP (Lightweight Directory Access Protocol), SSO (Single-Sign On), Kerberos or SAML (Security Assertion Markup Language) for identity management.

     b. In development or test environments, other Authentication Strategies may be employed.

Kubernetes does not have a notion of a human user.

/etc/kubernetes/admin.conf配置里边的证书信息可以通过下面方式解析
base64 -d | openssl x509 -noout -text

创建新用户流程：
首先，创建一个私钥文件（下面的“user.key”）
openssl genrsa -out user.key 2018
接下来，根据创建的私钥文件创建签名请求文件（CSR文件，以下称为“user.csr”）。您可以在这里输入国家名称、地区名称等，但由于 Kubernetes 只需要上面提到的组织名称和通用名称，因此您可以将其余部分留空
openssl req -new -key user.key -out user.csr

最后，使用 CA 证书和 CA 密钥文件从 CSR 文件创建证书文件（此处为“user.crt”）。
(这里的ca是k8s的root证书)
openssl x509 -req -in user.csr -CA /etc/kubernetes/pki/ca.crt -CAkey /etc/kubernetes/pki/ca.key -CAcreateserial
 -out user.crt -days 10000

要使用创建的证书对 API 服务器进行身份验证，请使用“kubectl config set-credentials <user name>”命令将要使用的证书信息添加到配置文件中。例如，要添加名为“testuser”的用户，请运行：

$ kubectl config set-credentials testuser --client-certificate=user.crt --client-key=user.key --embed-certs=true

↓ 查看设置的集群名称
$ kubectl config 获取集群
姓名
库伯内特斯

↓ 创建一个“test”上下文，将“kubernetes”集群与“testuser”用户关联起来
$ kubectl config set-context test --user=testuser --cluster=kubernetes
上下文“测试”已创建。

↓ 切换上下文以用于“测试”
$ kubectl config 使用上下文测试
切换到上下文“测试”。

↓查看上下文信息
$ kubectl 配置获取上下文
当前名称集群 AUTHINFO 命名空间
          kubernetes-admin@kubernetes kubernetes kubernetes-admin
* 测试 kubernetes testuser

### authentication
- oidc(https://www.youtube.com/watch?v=nPZ8QDZXtLI)

#### k8s运行动力
- 事件驱动,事件会堆积也会有优先级,不同的组件消费事件
#### scheduler
- watch pod crd(filter nodeName == "") schedule pod to node(更新pod cr的nodeName字段)
- 公平调度(不同事件的资源过来了需要及时处理)
- 资源高效利用(对node cpu memory资源的利用,cpu少了可能还能跑但是memory少了就不行了)
- QoS
- affinity
- anti-affinity
- data persistance(worker找数据)(马栏山:客户数据在哪个云上,就去哪个云上创建主机)
- predict(过滤)(filter chain narrow down候选节点)
-- hostport
-- resource
-- hostname
-- matchNodeSelector
-- affinity
-- diskConflict
-- tolerate taint
- priority(打分排序)
-- selectorSpreadPriority
-- interPodAffinity(topology zone)
-- leastRequestedPrior
-- balancedResourceAllocation
-- nodePreferAvoidPodsPrior
- resource(cpu)
-- limit:pod最多需要多少资源,调度器不考虑该值  1 -> 1000m 100m->0.1
-- request:调度器参考值,pod至少需要多少资源才能run(开发人员确定)
- resource(memory)
-- limit
-- request

#### initContainer
- 顺序执行,limit request取所有container中最大的那个就好了
- initContainer资源不会释放(pod restart 需要重新initContainer)
- 跟container不一样不是一个sum

如果不设置limit request会有什么问题,LimitRange可以设置默认limit request
#### request limit衍生的超卖
- pod通过配置较小的request run起来,当同node的其他pod负载很低的时候,该pod的资源使用在不超过limit的情况下是可以提高性能的

通过kubectl get node nodename 查看node status中的资源数量
- allocatable(可以分配给pod的资源总量)
- capbility
docker inspect 查看容器cgroup(cgroupParent,去节点上找到这个value对应的目录查看cgroup的值)
namespace做资源隔离,cgroup做资源限制

一个cpu period 100000微妙
创建的container的时候很耗资源(创建成功之后会降下来)
cpu set(绑定node某个cpu core)

#### controller-manager
- controller与controller之间的区别:工作流程相同,关注的对象不同,worker的逻辑不同
#### kubelet
#### CRI
#### CNI
#### CSI

#### list and watch(https://docs.bitnami.com/tutorials/a-deep-dive-into-kubernetes-controllers/)(https://github.com/kubernetes-sigs/controller-runtime/issues/521)
informer/reflect
shared informer(listwatch function/specific resource in specific namespace/eventhandler),listwatch里边有list watch函数,handler里边有add update delete函数,注意cache和workqueue的区别
work queue
case:
kubectl get resouce from apiserver 如果所有的get都要到达apiserver并查询etcd 压力会很大 于是有了缓存
缓存在哪里 如何让缓存跟etcd的数据保持一致


```go
lw := cache.NewListWatchFromClient(
      client,
      &v1.Pod{},
      api.NamespaceAll,
      fieldSelector)
store, controller := cache.NewInformer {
    &cache.ListWatch{},
    &v1.Pod{},
    resyncPeriod,
    cache.ResourceEventHandlerFuncs{},

```
event handler:
- AddFunc is called when a new resource is created.
- UpdateFunc is called when an existing resource is modified. The oldObj is the last known state of the resource. UpdateFunc is also called when a re-synchronization happens, and it gets called even if nothing changes.
- DeleteFunc is called when an existing resource is deleted. It gets the final state of the resource (if it is known). Otherwise, it gets an object of type DeletedFinalStateUnknown. This can happen if the watch is closed and misses the delete event and the controller doesn't notice the deletion until the subsequent re-list.

ResyncPeriod
ResyncPeriod defines how often the controller goes through all items remaining in the cache and fires the UpdateFunc again. This provides a kind of configuration to periodically verify the current state and make it like the desired state.

It's extremely useful in the case where the controller may have missed updates or prior actions failed. However, if you build a custom controller, you must be careful with the CPU load if the period time is too short.