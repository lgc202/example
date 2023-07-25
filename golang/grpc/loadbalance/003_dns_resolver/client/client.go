package main

import (
	greet "client/proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	addr := "server-service:8080"
	req := greet.Request{Name: "client"}
	conn, err := grpc.Dial(
		"dns:///"+addr,
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

// 实验步骤
// 1. 编译镜像
// docker build -t dns-resolver-client:v1 .
// docker build -t dns-resolver-server:v1 .

// 2. 部署server和client
// kubectl create ns dns-resolver
// kubectl apply -f server.yaml
// kubectl apply -f client.yaml

// 3. 验证 dns 解析到的是pod ip
// kubectl -n dns-resolver exec client-deployment-678b479c4f-lfvdv -it -- sh
// nslookup server-service.dns-resolver.svc.cluster.local

// 4. 查看 client 端 log 并进行扩缩容
// kubectl -n dns-resolver logs -f client-deployment-678b479c4f-lfvdv
// 当缩容后，由于 grpc 有连接探活机制，会自动丢弃（剔除）无效连接
// kubectl -n dns-resolver scale deployment server-deployment --replicas=2
// 当扩容后，由于没有感知机制，默认30秒才会去进行dns解析一次，新节点可能要一段时间后才能收到请求
// kubectl -n dns-resolver scale deployment server-deployment --replicas=3

// 5. 清理环境
// kubectl delete -f server.yaml
// kubectl delete -f server.yaml
// kubectl delete ns dns-resolver
