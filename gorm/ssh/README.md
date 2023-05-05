ssh 隧道方式远程连接DB

```shell
ssh -L 8888:127.0.0.1:3306 st-user@my-remote-server.host
```

1. 8888 随意指定本地端口
2. 127.0.0.1:3306 ssh 跳板机准备连接的服务
3. st-user@my-remote-server.host ssh的服务地址

