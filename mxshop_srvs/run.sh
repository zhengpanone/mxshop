# 编译 common.proto
python -m grpc_tools.protoc --mypy_out=./common/proto/pb --python_out=./common/proto/pb --grpc_python_out=./common/proto/pb -I=common/proto ./common/proto/common.proto ./common/proto/user.proto
python -m grpc_tools.protoc --mypy_out=./common/proto/pb --python_out=./common/proto/pb --grpc_python_out=./common/proto/pb -I=common/proto ./common/proto/common.proto ./common/proto/role.proto
python -m grpc_tools.protoc --mypy_out=./common/proto/pb --python_out=./common/proto/pb --grpc_python_out=./common/proto/pb -I=common/proto ./common/proto/common.proto ./common/proto/dict.proto



