package _02_etcd_resolver

import (
	greet "002_etcd_resolver/proto"
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	etcdResolver "go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"log"
	"net"
	"strings"
	"testing"
	"time"
)

func TestEtcdResolver(t *testing.T) {
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

	// 启动服务
	server := grpc.NewServer()
	greet.RegisterGreeterServer(server, &greeterServer{addr: addr})
	go func() {
		if err = server.Serve(listener); err != nil {
			log.Fatalf(err.Error())
		}
	}()

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
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
		// 取消注册, 如果不主动撤销, 客户端还会继续往服务器发送请求, 直到租约过期
		if err := em.DeleteEndpoint(client.Ctx(), ServiceName+"/"+addr); err != nil {
			log.Println(err)
		}
		log.Printf("server: %s exit", addr)

		_ = client.Close()
	}()

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

func StartClient() {
	// 创建并注册etcdResolver
	c, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
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
	builder, err := etcdResolver.NewBuilder(c)
	if err != nil {
		return err
	}

	resolver.Register(builder)
	return nil
}
