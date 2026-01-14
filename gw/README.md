# Mundo Gateway SDK

网关 SDK，用于服务注册和路由管理。

## 从 mundo-gateway-sdk 迁移说明

该包已从独立的 `mundo-gateway-sdk` 模块整合到 `mundo-proto-sdk/gateway` 中。

### 导入路径变更

**旧的导入方式：**
```go
import gatewaySdk "github.com/trancecho/mundo-gateway-sdk"
```

**新的导入方式：**
```go
import gatewaySdk "github.com/trancecho/mundo-proto-sdk/gateway"
```

## 使用示例

### HTTP 服务注册

```go
import (
    "github.com/gin-gonic/gin"
    gateway "github.com/trancecho/mundo-proto-sdk/gateway"
)

func main() {
    r := gin.Default()
    
    gw := gateway.NewGatewayService(
        "your-service-name",
        "your-service-address",
        "http",
        "gateway-url",
        gateway.NewMyRedisTokenGetter("redis-addr", "password", 0),
    )
    
    gw.HttpConn(r)
}
```

### gRPC 服务注册

```go
import (
    gateway "github.com/trancecho/mundo-proto-sdk/gateway"
    "google.golang.org/grpc"
)

func main() {
    server := grpc.NewServer()
    
    gw := gateway.NewGatewayService(
        "your-service-name",
        "your-service-address",
        "grpc",
        "gateway-url",
        gateway.NewMyRedisTokenGetter("redis-addr", "password", 0),
    )
    
    gw.GrpcConn(server)
}
```

## 主要功能

- **服务注册**: 自动注册服务到网关
- **心跳保活**: 定期发送心跳信号保持服务活跃
- **路由注册**: 自动注册 HTTP/gRPC 路由到网关
- **客户端**: 支持通过网关获取目标服务地址

## 结构说明

- `sdk.go` - 网关服务主要实现
- `client.go` - 网关客户端，用于获取目标服务
- `get_secret.go` - Redis Token 获取器
- `model/gw.go` - 数据模型定义
- `i/gatewayv2.go` - 接口定义
