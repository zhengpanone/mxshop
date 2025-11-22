# 初始化
```shell
uv venv --python 3.12
uv sync
```

# 构建镜像
```shell
docker build --no-cache -f goods-web/Dockerfile -t zhengpanone/mxshop_srvs/goods_srv .
docker build --no-cache -f order-web/Dockerfile -t zhengpanone/mxshop-api/order-web .
docker build --no-cache -f user-web/Dockerfile -t zhengpanone/mxshop-api/user-web .
docker build --no-cache -f userop_srv/Dockerfile -t zhengpanone/mxshop_srvs/userop_srv .
```