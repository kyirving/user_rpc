### 用户RPC服务
使用框架 `go-zero`
生成RPC代码
```bash
mkdir user_rpc && cd user_rpc

# 获取公共协议
go get github.com/kyirving/common_proto@v1.0.0

# 生成RPC代码
PROTO_PATH=$(go list -m -f {{.Dir}} github.com/kyirving/common_proto)
goctl rpc protoc $PROTO_PATH/user/user.proto \
  --go_out=. --go-grpc_out=. --zrpc_out=. \
  --client=false \
  -I $PROTO_PATH

```