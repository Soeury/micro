package createresolver

import "google.golang.org/grpc/resolver"

// 步骤：
//   -1. 重写 resolver.Resolver 接口(定义一个结构体，并实现 ResolveNow 和 Close 方法)

//        type Resolver interface {
//	          ResolveNow(ResolveNowOptions)
//	          Close()
//        }

//   -2. 重写 Builder 接口(定义一个结构体，并实现 Build 和 Scheme 方法)

//       type Builder interface {
//	         Build(target Target, cc ClientConn, opts BuildOptions) (Resolver, error)
//	         Scheme() string
//       }

const (
	myScheme   = "rabbit"
	myEndpoint = "resolver.kedudu.com"
)

var addrs = []string{"127.0.0.1:8080", "127.0.0.1:8081", "127.0.0.1:8082"}

// 重写了 resolver.Resolver 接口
type RabbitResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string][]string
}

func (r *RabbitResolver) ResolveNow(o resolver.ResolveNowOptions) {

	addrStrs := r.addrsStore[r.target.Endpoint()]
	addrList := make([]resolver.Address, len(addrStrs))

	for idx, s := range addrStrs {
		addrList[idx] = resolver.Address{Addr: s}
	}
	r.cc.UpdateState(resolver.State{Addresses: addrList})
}

func (*RabbitResolver) Close() {}

// q1miResolverBuilder 需实现 Builder 接口
type RabbitResolverBuilder struct{}

func (*RabbitResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {

	r := &RabbitResolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]string{
			myEndpoint: addrs,
		},
	}

	r.ResolveNow(resolver.ResolveNowOptions{})

	// 这里 RabbitResolver 重写了 resolver.Resolver 接口，所以可以使用 resolver.Resolver 作为返回值类型
	return r, nil
}

func (*RabbitResolverBuilder) Scheme() string {

	return myScheme
}

func Init() {

	resolver.Register(&RabbitResolverBuilder{})
}
