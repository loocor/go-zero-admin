# 初始化流程

## 1、安装

- 框架安装: `go get github.com/tal-tech/go-zero`
- 框架代码生成工具安装:  `go get -u github.com/tal-tech/go-zero/tools/goctl`

## 2、创建 API

-  进到 `api/` 目录执行: 

```bash
goctl api go -api ./doc/admin.api  -dir .
```

## 3、创建 RPC

- 进到 `rpc/sys/` 目录操作: 

```bash
goctl rpc proto -src sys.proto -dir .
```

- 进到 `rpc/ums/` 目录操作:

```bash
goctl rpc proto -src ums.proto -dir .
```

## 4、创建 Model

进到 `model/` 目录操作:

```bash
goctl model mysql ddl -c -src book.sql -dir .

goctl model mysql datasource -url="root:password@tcp(127.0.0.1:3306)/gozero" -table="sys*" -dir ./model/sys
goctl model mysql datasource -url="root:password@tcp(127.0.0.1:3306)/gozero" -table="ums*" -dir ./model/ums
```
