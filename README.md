# GO-API
Golang api framework.

## Getting Started | 快速开始
### Requirements | 环境要求

 - Golang > 1.19
 - Mysql 或 PostgreSQL
 - Redis（可选）

Golang 配置：

```sh
$ go env -w GOPROXY=https://goproxy.cn,direct
$ go env -w GO111MODULE=on
```

### Installation | 安装

clone：

```sh
$ git clone https://github.com/johncxf/laradmin.git
```

安装依赖：

```sh
$ go mod tidy
```

### Usage | 用法

修改配置文件：

```sh
# 新建配置文件
$ mv config/.env.yml config/env.yml

# 按照需求修改配置文件
$ vim config/env.yml
```

直接启动：

```sh
$ go run main.go
```

或编译运行：

```sh
# 编译为二进制文件
$ go build -o ./bin/go-api .

# 运行
$ ./bin/go-api config/env.yml
```

或使用docker部署启动：

```sh
$ docker-compose up -d
```

## Contributing | 贡献

## FAQ