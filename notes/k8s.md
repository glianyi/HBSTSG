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