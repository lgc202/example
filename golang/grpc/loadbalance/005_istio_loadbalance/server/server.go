package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	greet "server/proto"
)

type greeterServer struct {
	greet.UnimplementedGreeterServer
}

func (g *greeterServer) Greet(ctx context.Context, request *greet.Request) (*greet.Response, error) {
	name, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	return &greet.Response{
		Greet: name,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf(err.Error())
	}

	// 启动服务
	server := grpc.NewServer()
	greet.RegisterGreeterServer(server, &greeterServer{})
	if err = server.Serve(listener); err != nil {
		log.Fatalf(err.Error())
	}
}
