goctl api go -api user.api -dir .

goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.

goctl model mysql datasource -url="root:root@tcp(127.0.0.1:3306)/order" -table="*"  -dir="./model" -c
