# 构建阶段 - 使用 golang 镜像编译三个服务
FROM golang:latest AS builder

# 安装必要的工具
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates tzdata && rm -rf /var/lib/apt/lists/*

# 设置 Go 代理加速依赖下载
ENV GOPROXY=https://goproxy.cn,direct

# 设置工作目录
WORKDIR /build

# 复制 go.mod 和 go.sum
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制整个项目源码
COPY . .

# 编译三个服务
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -tags no_k8s -o ./bin-chinese ./chinese/chinese.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -tags no_k8s -o ./bin-english ./english/english.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -tags no_k8s -o ./bin-greet ./greet/greet.go

# 运行阶段 - 使用 Alpine 作为基础镜像
FROM alpine:latest

# 安装必要的运行时依赖
RUN apk add --no-cache ca-certificates tzdata wget

# 创建非 root 用户
RUN addgroup -g 1000 app && \
    adduser -D -u 1000 -G app app

# 设置时区（可根据需要调整）
ENV TZ=Asia/Shanghai

# 创建日志目录
RUN mkdir -p /app/logs && \
    chown -R app:app /app

# 创建各个服务的目录结构
RUN mkdir -p /app/chinese/etc /app/chinese/logs && \
    mkdir -p /app/english/etc /app/english/logs && \
    mkdir -p /app/greet/etc /app/greet/logs && \
    chown -R app:app /app

# 从构建阶段复制编译好的二进制文件
COPY --from=builder /build/bin-chinese /app/chinese/chinese
COPY --from=builder /build/bin-english /app/english/english
COPY --from=builder /build/bin-greet /app/greet/greet

# 复制配置文件
COPY --chown=app:app chinese/etc/chinese.yaml /app/chinese/etc/
COPY --chown=app:app english/etc/english.yaml /app/english/etc/
COPY --chown=app:app greet/etc/greet-api.yaml /app/greet/etc/

# 设置工作目录
WORKDIR /app

# 切换到非 root 用户
USER app

# 默认命令（实际由 docker-compose.yml 覆盖）
CMD ["/bin/sh", "-c", "echo 'Use docker-compose to start specific services'"]
