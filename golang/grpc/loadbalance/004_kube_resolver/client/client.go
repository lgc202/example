package main

import (
	greet "client/proto"
	"context"
	"github.com/zeromicro/go-zero/zrpc/resolver"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	// 注册 go-zero 实现的 resolver
	resolver.Register()

	// kube-resolver 是服务所在的namespace
	addr := "server-service:8080"
	req := greet.Request{Name: "client"}
	conn, err := grpc.Dial(
		"k8s://kube-resolver/"+addr,
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
// docker build -t kube-resolver-client:v1 .
// docker build -t kube-resolver-server:v1 .
// 如果 k8s 环境的container runtime是containerd, 还需要转换一下镜像
// docker save kube-resolver-client:v1 -o kube-resolver-client.tar
// docker save kube-resolver-server:v1 -o kube-resolver-server.tar
// export ctr="ctr -a /run/k3s/containerd/containerd.sock --namespace k8s.io"
// ctr image import kube-resolver-client.tar && rm kube-resolver-client.tar
// ctr image import kube-resolver-server.tar && rm kube-resolver-server.tar
// 查看是否成功: crictl images | grep kube-resolver

// 2. 部署server和client
// kubectl create ns kube-resolver
// kubectl apply -f server.yaml
// kubectl apply -f service-account.yaml
// kubectl apply -f client.yaml

// 3. 验证 dns 解析到的是服务器的service ip
// kubectl -n kube-resolver exec client-deployment-678b479c4f-lfvdv -it -- sh
// nslookup server-service.kube-resolver.svc.cluster.local

// 4. 查看 client 端 log 并进行扩缩容
// kubectl -n kube-resolver logs -f client-deployment-678b479c4f-lfvdv
// 当缩容后，客户端可以watch到endpoint的修改，把无效连接丢弃掉
// kubectl -n kube-resolver scale deployment server-deployment --replicas=2
// 当扩容后，客户端可以watch到endpoint的修改，新节点很快可以接收到请求
// kubectl -n kube-resolver scale deployment server-deployment --replicas=3

// 5. 清理环境
// kubectl delete -f server.yaml
// kubectl delete -f service-account.yaml
// kubectl delete -f client.yaml
// kubectl delete ns kube-resolver
