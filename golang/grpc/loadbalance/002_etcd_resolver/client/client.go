package main

import (
	greet "002_etcd_resolver/proto"
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	etcdResolver "go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"log"
	"time"
)

func main() {
	// 创建并注册etcdResolver
	c, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2382"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	if err := registerEtcdResolver(c); err != nil {
		log.Fatalf(err.Error())
	}

	// 建立对应 scheme 的连接, 并且配置负载均衡
	req := greet.Request{Name: "client"}
	conn, err := grpc.Dial(
		fmt.Sprintf("etcd:///%s", ServiceName),
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
		time.Sleep(time.Second)
		response, err := client.Greet(context.Background(), &req)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		log.Printf("got the response from server: %s\n", response.GetGreet())
	}
}

const (
	ServiceName = "greet"
)

func registerEtcdResolver(c *clientv3.Client) error {
	// 如果遇到以下错误, 就降低一下 grpc 的版本
	// cannot use target.Endpoint (value of type func() string) as string value in struct literal
	// 比如用v1.52.3: go get google.golang.org/grpc@v1.52.3
	builder, err := etcdResolver.NewBuilder(c)
	if err != nil {
		return err
	}

	resolver.Register(builder)
	return nil
}
