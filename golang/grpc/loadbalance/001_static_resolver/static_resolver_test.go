package _01_static_resolver

import (
	greet "001_static_resolver/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"log"
	"net"
	"testing"
	"time"
)

func TestResolveStaticAddr(t *testing.T) {
	go StartServer("localhost:8080")
	go StartServer("localhost:8081")

	time.Sleep(time.Second * 3)

	StartClient()
}

type greeterServer struct {
	greet.UnimplementedGreeterServer
	addr string
}

func (g *greeterServer) Greet(context.Context, *greet.Request) (*greet.Response, error) {
	return &greet.Response{Greet: g.addr}, nil
}

func StartServer(addr string) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf(err.Error())
	}

	server := grpc.NewServer()
	greet.RegisterGreeterServer(server, &greeterServer{addr: addr})
	if err = server.Serve(listener); err != nil {
		log.Fatalf(err.Error())
	}
}

func StartClient() {
	// 注册我们的 resolver
	resolver.Register(NewExampleResolverBuilder())

	// 建立对应 scheme 的连接, 并且配置负载均衡
	// example:///test 中的 test 一般是远程服务的服务名, 如果使用etcd等作为注册中心，可用该服务名去注册中心找出
	// 属于该服务的所有实例，在这个例子中 test 的值可以随便填
	req := greet.Request{Name: "client"}
	conn, err := grpc.Dial(
		"example:///test",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)

	if err != nil {
		log.Fatalf(err.Error())
	}
	defer func() {
		_ = conn.Close()
	}()

	client := greet.NewGreeterClient(conn)
	for {
		response, err := client.Greet(context.Background(), &req)
		if err != nil {
			log.Fatalf(err.Error())
		}

		log.Printf("got the response from server: %s\n", response.GetGreet())
		time.Sleep(time.Second)
	}
}
