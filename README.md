# Database Course Design服务器端代码

## 先决条件

若在服务器上部署此代码，则需满足以下先决条件：

1. Oracle数据库以及与目标Oracle数据库兼容的Oracle Instant Client
2. Go语言运行时

注：请考虑Oracle Instant Client版本与Oracle Database的兼容性。测试时发现Oracle 11gR2与Instant Client>=19.3不兼容。

## 部署

@TODO 补充部署过程

1. `go get <github-repo-url>`
2. 将`config.go.example`重命名为`config.go`，并根据实际情况修改文件中的信息。

### 使用Nginx设置反向代理

```conf
server {
    server_name  dbcd.qfstudio.net;

    location / {
        proxy_pass http://localhost:12323;
    }
}
```

## 启动和终止

```sh
$ netstat -ap | grep 12323
tcp        0      0 127.0.0.1:12323         0.0.0.0:*               LISTEN      4639/dbcd
$ kill 4639
```
