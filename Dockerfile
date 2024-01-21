FROM golang:alpine

# 在容器内部设置环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 设置后续指令的工作目录
WORKDIR /build

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .

# 将代码编译成二进制可执行文件
RUN go build -o go-api .

WORKDIR /dist

RUN cp /build/go-api .

COPY ./config ./config

# 需要运行的命令
ENTRYPOINT ["/dist/go-api", "/dist/config/env.yml"]