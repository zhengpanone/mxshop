# swaggo
```shell
swag init -g ./user-web/cmd/main.go -d . --exclude golang.org/x/exp --parseDependency true
swag init -g ./goods-web/cmd/main.go -d .
swag fmt
cd userop-web && swag init -g cmd/main.go -d . --parseDependency=true --parseInternal=true

swag init -g cmd/main.go -d ./api/controller --parseDependency --parseDepth=6 -o ./docs
```

# 构建镜像
```shell
docker build --no-cache -f goods-web/Dockerfile -t zhengpanone/mxshop-api/goods-web .
docker build --no-cache -f order-web/Dockerfile -t zhengpanone/mxshop-api/order-web .
docker build --no-cache -f user-web/Dockerfile -t zhengpanone/mxshop-api/user-web .
docker build --no-cache -f userop-web/Dockerfile -t zhengpanone/mxshop-api/userop-web .
```