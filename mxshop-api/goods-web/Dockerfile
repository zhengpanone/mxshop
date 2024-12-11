#https://mp.weixin.qq.com/s/vGjvIb4uWRCbuzafZuljyw
# 使用 Go 官方的 alpine 镜像作为基础镜像
FROM golang:alpine AS mod
# 设置 Go 代理
ENV GOPROXY=https://goproxy.cn,direct
# 安装 git 依赖
RUN apk add --no-cache git
# 设置工作目录为 /workspace
WORKDIR /workspace/mxshop-api

# 将上层目录的所有内容复制到 /workspace
COPY ../common /workspace/mxshop-api/common
COPY . .

# 切换到 goods-web 目录
WORKDIR /workspace/mxshop-api/goods-web


# 下载 Go 依赖
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod tidy && go mod download


# 使用 mod 阶段的结果，继续在同一目录下构建应用程序
FROM mod AS build
# 在 /workspace 目录中进行构建
RUN  CGO_ENABLED=0 go build -o goods-web -ldflags '-w -extldflags "-static"' .


# 使用 Alpine 作为最终镜像，并将构建的应用程序复制进去
FROM alpine:latest
# 安装时区数据和 CA 证书
RUN apk add --no-cache tzdata ca-certificates
# 将构建好的应用程序从 build 阶段复制到最终镜像中
WORKDIR /app
COPY --from=build /workspace/mxshop-api/goods-web/goods-web /app
COPY --from=build /workspace/mxshop-api/goods-web/config-dev.yaml /app

# 设置容器启动时运行的命令
CMD ["/app/goods-web"]