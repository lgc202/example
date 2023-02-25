# go-zero 的相关练习

## 1. 环境搭建
```shell
# 安装 Goctl, go1.16+
$ go install github.com/zeromicro/go-zero/tools/goctl@latest
# 安装 protobuf
$ goctl env check -i -f --verbose
```
## 2. 单体服务
(1) 生成项目框架
```shell
$ mkdir single-demo
$ cd single-demo
$ go mod init single-demo
$ goctl api new greet
$ go mod tidy
```
生成的目录结构如下
```shell
.
├── go.mod
├── go.sum
└── greet
    ├── etc
    │   └── greet-api.yaml
    ├── greet.api
    ├── greet.go
    └── internal
        ├── config
        │   └── config.go
        ├── handler
        │   ├── greethandler.go
        │   └── routes.go
        ├── logic
        │   └── greetlogic.go
        ├── svc
        │   └── servicecontext.go
        └── types
            └── types.go
```
（2）在greetlogic.go编写业务逻辑
```go
func (l *GreetLogic) Greet(req *types.Request) (resp *types.Response, err error) {
	return &types.Response{
		Message: "hello world",
	}, nil
}
```
（3）启动服务并测试
```shell
# 启动服务
$ cd greet
$ go run greet.go -f etc/greet-api.yaml
Starting server at 0.0.0.0:8888...

# 测试
$ curl -i -X GET http://localhost:8888/from/you
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Traceparent: 00-30190724ed5b81f0bd2a77893bb67016-e50914440a8571b7-00
Date: Mon, 09 Jan 2023 02:44:15 GMT
Content-Length: 25

{"message":"hello world"}#
```
## 3. 微服务
该例子有两个微服务  
- 订单服务(order)提供一个查询接口  
- 用户服务(user)提供一个方法供订单服务获取用户信息

（1）生成项目框架
```shell
$ mkdir micro-demo
$ cd micro-demo
$ go mod init micro-demo
```
（2）创建 user rpc 服务  
```shell
$ mkdir -p mall/user/rpc
```
添加user.proto文件，增加getUser方法
```protobuf
syntax = "proto3";

package user;
  
// protoc-gen-go 版本大于1.4.0, proto文件需要加上go_package,否则无法生成
option go_package = "./user";

message IdRequest {
    string id = 1;
}
  
message UserResponse {
    // 用户id
    string id = 1;
    // 用户名称
    string name = 2;
    // 用户性别
    string gender = 3;
}
  
service User {
    rpc getUser(IdRequest) returns(UserResponse);
}
```
生成代码
```shell
$ cd mall/user/rpc
$ goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
$ go mod tidy
```
在getuserlogic中添加业务逻辑
```go
func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
	return &user.UserResponse{
		Id:   "1",
		Name: "test",
	}, nil
}
```
（3）创建order api服务
```shell
# 回到micro-demo目录
$ mkdir -p order/api && cd order/api
```
添加order.api文件  
```shell
type(
    OrderReq {
        Id string `path:"id"`
    }

    OrderReply {
        Id string `json:"id"`
        Name string `json:"name"`
    }
)
service order {
    @handler getOrder
    get /api/order/get/:id (OrderReq) returns (OrderReply)
}
```
生成order服务
```shell
$ goctl api go -api order.api -dir .
$ go mod tidy
```
添加user rpc配置
- 修改config.go
```go
type Config struct {
    rest.RestConf
    UserRpc zrpc.RpcClientConf
}
```
- 修改yaml配置
```yaml
Name: order
Host: 0.0.0.0
Port: 8888
UserRpc:
  Etcd:
    Hosts:
    - 127.0.0.1:2379
    Key: user.rpc
```
完善服务依赖, 修改servicecontext.go
```go
type ServiceContext struct {
    Config  config.Config
    UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
    return &ServiceContext{
        Config:  c,
        UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
    }
}
```
添加order演示逻辑, 修改getorderlogic 
```go

func (l *GetOrderLogic) GetOrder(req *types.OrderReq) (*types.OrderReply, error) {
    user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{
        Id: "1",
    })
    if err != nil {
        return nil, err
    }

    if user.Name != "test" {
        return nil, errors.New("用户不存在")
    }

    return &types.OrderReply{
        Id:   req.Id,
        Name: "test order",
    }, nil
}
```
启动服务并验证
```shell
# 启动etcd
$ etcd
或指定端口
$ etcd --listen-client-urls 'http://localhost:2382' --advertise-client-urls 'http://localhost:2382'

# 启动user rpc
go run user.go -f etc/user.yaml

# 启动order api
$ go run order.go -f etc/order.yaml

# 访问order api
$ curl -i -X GET http://localhost:8888/api/order/get/1
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Traceparent: 00-9f695afb9fc0feb4c1abd673db27a2f4-2e4b2ef44cd8f15e-00
Date: Mon, 09 Jan 2023 08:07:59 GMT
Content-Length: 30

{"id":"1","name":"test order"}#
```