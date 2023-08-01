package main

import (
	greet "001_static_resolver/proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
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

	server := grpc.NewServer()
	greet.RegisterGreeterServer(server, &greeterServer{addr: addr})
	if err = server.Serve(listener); err != nil {
		log.Fatalf(err.Error())
	}
}

func main() {
	go StartServer("localhost:8080")
	go StartServer("localhost:8081")

	select {}
}
