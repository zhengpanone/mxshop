#https://mp.weixin.qq.com/s/vGjvIb4uWRCbuzafZuljyw
# 使用 Go 官方的 Alpine 镜像作为基础镜像
FROM golang:alpine AS mod
# 设置 Go 代理，并安装 git 和构建所需的依赖
ENV GOPROXY=https://goproxy.cn,direct
RUN apk add --no-cache git
# 设置工作目录
WORKDIR /workspace/mxshop-api

# 将上层目录的所有内容复制到容器中
COPY ../common /workspace/mxshop-api/common

# 切换到 order-web 目录
WORKDIR /workspace/mxshop-api/order-web


# 下载 Go 依赖
RUN go env -w GOPROXY=https://goproxy.cn,direct && \
    go mod tidy && go mod download

# 使用 mod 阶段的结果，继续构建应用
FROM mod AS build
# 构建 Go 应用
RUN CGO_ENABLED=0 go build -o order-web -ldflags '-w -extldflags "-static"' .

# 使用更小的 Alpine 镜像作为最终镜像
FROM alpine:latest

# 安装时区数据和 CA 证书（由第一阶段的构建继承）
RUN apk add --no-cache tzdata ca-certificates

# 设置工作目录
WORKDIR /app

# 复制构建好的应用程序和配置文件
COPY --from=build /workspace/mxshop-api/order-web/order-web /app
COPY --from=build /workspace/mxshop-api/order-web/config-dev.yaml /app

# 设置容器启动时运行的命令
CMD ["/app/order-web"]