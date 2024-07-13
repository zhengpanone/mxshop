protoc -I . order.proto --go_out=plugins=grpc:.
protoc -I . goods.proto --go_out=plugins=grpc:.
protoc -I . inventory.proto --go_out=plugins=grpc:.