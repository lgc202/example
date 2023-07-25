package _01_static_resolver

import (
	"google.golang.org/grpc/resolver"
)

const (
	exampleScheme = "example"
)

type exampleResolverBuilder struct {
}

func NewExampleResolverBuilder() *exampleResolverBuilder {
	return &exampleResolverBuilder{}
}

func (e *exampleResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &exampleResolver{}

	addrStr := []string{"localhost:8080", "localhost:8081"}
	addr := make([]resolver.Address, len(addrStr))
	for _, s := range addrStr {
		addr = append(addr, resolver.Address{Addr: s})
	}

	if err := cc.UpdateState(resolver.State{Addresses: addr}); err != nil {
		cc.ReportError(err)
	}

	return r, nil
}

func (e *exampleResolverBuilder) Scheme() string {
	return exampleScheme
}

type exampleResolver struct{}

func (e *exampleResolver) ResolveNow(options resolver.ResolveNowOptions) {}

func (e *exampleResolver) Close() {}
