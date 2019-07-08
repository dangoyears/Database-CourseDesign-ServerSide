# Database Course Design服务器端代码

[![Build Status](https://travis-ci.org/dangoyears/Database-CourseDesign-ServerSide.svg?branch=master)](https://travis-ci.org/dangoyears/Database-CourseDesign-ServerSide)
[![Go Report Card](https://goreportcard.com/badge/github.com/dangoyears/Database-CourseDesign-ServerSide)](https://goreportcard.com/report/github.com/dangoyears/Database-CourseDesign-ServerSide)

## 先决条件

若在服务器上部署此代码，则需满足以下先决条件：

1. Oracle数据库以及与目标Oracle数据库兼容的Oracle Instant Client
2. Go语言运行时

注：请考虑Oracle Instant Client版本与Oracle Database的兼容性。测试时发现Oracle 11gR2与Instant Client>=19.3不兼容。

## 部署

1. `go get github.com/dangoyears/Database-CourseDesign-ServerSide`
2. 将`config.example.yaml`重命名为`config.yaml`并正确配置。

## 其他提示

### 使用Nginx设置反向代理

```conf
server {
    server_name  dbcd.qfstudio.net;

    location / {
        proxy_pass http://localhost:12323;
    }
}
```

### 启动和终止

```sh
$ netstat -ap | grep 12323
tcp        0      0 127.0.0.1:12323         0.0.0.0:*               LISTEN      4639/dbcd
$ kill 4639
```

### 关于Oracle SQL的占位符

对于形如`update "Human" set "Name"=:2 where "HumanID"=:1`等带绑定的SQL语句，goracle.v2`3222d7159b45fce95150f06a57e1bcc2868108d3`不会按照SQL语句中占位符的数字所暗示的参数顺序进行绑定，而是按照占位符出现的顺序进行绑定。
