FROM golang:1.21.4-alpine3.18 AS builder

# 设置工作目录
WORKDIR /app

# 复制应用程序代码到容器中
COPY . .

# 下载依赖项
RUN go mod download

# 构建应用程序
RUN go build -o main .

# 使用轻量级基础镜像
FROM alpine:3.18

# 设置工作目录
WORKDIR /app

# 复制应用程序到容器中
COPY --from=builder /app/main .

# 暴露端口
EXPOSE 9000

# 启动应用程序
ENTRYPOINT ["./main"]
