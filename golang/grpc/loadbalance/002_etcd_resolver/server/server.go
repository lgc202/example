package main

import (
	greet "002_etcd_resolver/proto"
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
	"time"
)

const (
	ServiceName = "greet"
)

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

	// 启动服务
	server := grpc.NewServer()
	greet.RegisterGreeterServer(server, &greeterServer{addr: addr})
	go func() {
		if err = server.Serve(listener); err != nil {
			log.Fatalf(err.Error())
		}
	}()

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2382"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	em, err := endpoints.NewManager(client, ServiceName)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// 创建租约，如果不续约，10 秒后会过期被删除
	lease := clientv3.NewLease(client)
	resp, err := lease.Grant(context.Background(), 10)
	if err != nil {
		log.Fatal(err)
	}

	// 将服务注册到 etcd
	err = em.AddEndpoint(client.Ctx(), ServiceName+"/"+addr, endpoints.Endpoint{Addr: addr}, clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		// 取消注册
		if err := em.DeleteEndpoint(client.Ctx(), ServiceName+"/"+addr); err != nil {
			log.Println(err)
		}
		log.Printf("server: %s exit", addr)

		_ = client.Close()
	}()

	log.Printf("start server success on: %s\n", addr)
	timer := time.NewTicker(30 * time.Second)
	for {
		select {
		case <-time.After(5 * time.Second):
			// 每隔 5 s进行一次延续租约的动作
			if _, err := client.KeepAliveOnce(client.Ctx(), resp.ID); err != nil {
				log.Fatalf(err.Error())
			}
		case <-timer.C:
			// 1 分钟后退出 8080 的服务器，8081继续运行
			if strings.Contains(addr, "8080") {
				return
			}
			timer.Stop()
		case <-client.Ctx().Done():
			return
		}
	}
}

func main() {
	go StartServer("localhost:8080")
	go StartServer("localhost:8081")

	select {}
}
