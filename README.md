# kobe

经过修改，在CreateProject时，不去拉取代码仓库，改为创inventory yaml文件

kobe 是 Ansible 的轻量级封装，提供了 grpc api 和 golang cli。

## 主要功能

- playbook
- adhoc

## build
```shell
make generate_grpc

make docker
```

## 本地测试
docker run --rm -p 8082:8080 --name kobe rubinus/kobe:v1.0


 
