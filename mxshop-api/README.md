# swaggo
```shell
swag init -g ./user-web/cmd/main.go -d . --exclude golang.org/x/exp --parseDependency true
swag init -g ./goods-web/cmd/main.go -d .
swag fmt
cd userop-web && swag init -g cmd/main.go -d . --parseDependency=true --parseInternal=true

```

# 构建镜像
```shell
docker build --no-cache -f goods-web/Dockerfile -t zhengpanone/goods-web .
```