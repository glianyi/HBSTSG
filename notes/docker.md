docker rmi $(sudo docker images -f "dangling=true" -q)

### 网络