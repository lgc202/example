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
		addr,
		grpc.WithInsecure(),
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
// 1. 安装 istio(最新版已经到1.18.2,但我的k8s版本不支持)
// curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.13.1 TARGET_ARCH=x86_64 sh -
// export PATH="$PATH:/usr/local/src/istio-1.13.1/bin"
// istioctl install --set profile=demo -y
// 查看 是否安装成功
// kubectl get po,svc -n istio-system
// 如果你的环境不支持loadbalance类型的service, 可改成NodePort类型的
// kubectl patch svc -n istio-system istio-ingressgateway -p '{"spec": {"type": "NodePort"}}'
// 如要卸载, 执行以下命令
// istioctl manifest generate --set profile=demo | kubectl delete --ignore-not-found=true -f -

// 2. 编译镜像
// docker build -t istio-loadbalance-client:v1 .
// docker build -t istio-loadbalance-server:v1 .
// 如果 k8s 环境的container runtime是containerd, 还需要转换一下镜像
// docker save istio-loadbalance-client:v1 -o istio-loadbalance-client.tar
// docker save istio-loadbalance-server:v1 -o istio-loadbalance-server.tar
// export ctr="ctr -a /run/k3s/containerd/containerd.sock --namespace k8s.io"
// ctr image import istio-loadbalance-client.tar && rm istio-loadbalance-client.tar
// ctr image import istio-loadbalance-server.tar && rm istio-loadbalance-server.tar
// 查看是否成功: crictl images | grep istio-loadbalance

// 3. 部署server和client
// kubectl create ns istio-loadbalance
// 给 namespace 打上 label, 允许注入 istio
// kubectl label namespace istio-loadbalance istio-injection=enabled --overwrite
// kubectl apply -f server.yaml
// kubectl apply -f client.yaml
// apply完成可以看到pod里面有两个容器, 一个是sidecar另一个是业务容器

// 4. 查看 client 端 log 并进行扩缩容
// kubectl -n istio-loadbalance logs -f client-deployment-678b479c4f-lfvdv
// 缩容
// kubectl -n istio-loadbalance scale deployment server-deployment --replicas=2
// 扩容
// kubectl -n istio-loadbalance scale deployment server-deployment --replicas=3

// 5. 清理环境
// kubectl delete -f server.yaml
// kubectl delete -f client.yaml
// kubectl delete ns istio-loadbalance
