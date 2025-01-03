# swaggo
```shell
swag init -g ./user-web/cmd/main.go -d . --exclude golang.org/x/exp --parseDependency true
swag init -g ./goods-web/cmd/main.go -d .
swag fmt
```

# 构建镜像
```shell
docker build --no-cache -f goods-web/Dockerfile -t goods-web .
```