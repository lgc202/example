package main

import (
	staticresolver "001_static_resolver"
	greet "001_static_resolver/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"log"
	"time"
)

func main() {
	// 注册我们的 resolver
	resolver.Register(staticresolver.NewExampleResolverBuilder())

	// 建立对应 scheme 的连接, 并且配置负载均衡
	// example:///test 中的 test 一般是远程服务的服务名, 如果使用etcd等作为注册中心，可用该服务名去注册中心找出
	// 属于该服务的所有实例，在这个例子中 test 的值可以随便填
	req := greet.Request{Name: "client"}
	conn, err := grpc.Dial(
		//"example:///test",
		"localhost:8080",
		grpc.WithInsecure(),
		//grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
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
			log.Println(err.Error())
			log.Println(conn.GetState())
			continue
		}

		log.Printf("got the response from server: %s\n", response.GetGreet())
		time.Sleep(time.Second)
	}
}
