package resolve

import (
	"context"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
)

// etcdResolverBuilder 实现了 resolver.Builder 接口
type EtcdResolverBuilder struct {
	EtcdClient *clientv3.Client
}

func (b *EtcdResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &etcdResolver{
		target:     target,
		cc:         cc,
		etcdClient: b.EtcdClient,
	}
	go r.watch()
	return r, nil
}

func (b *EtcdResolverBuilder) Scheme() string {
	return "etcd"
}

// etcdResolver 实现了 resolver.Resolver 接口
type etcdResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	etcdClient *clientv3.Client
}

func (r *etcdResolver) ResolveNow(o resolver.ResolveNowOptions) {
	r.watch()
}

func (r *etcdResolver) Close() {}

func (r *etcdResolver) watch() {
for {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		resp, err := r.etcdClient.Get(ctx, "services/"+r.target.Endpoint(), clientv3.WithPrefix())
		cancel()
		if err != nil {
			r.cc.ReportError(err)
			time.Sleep(2 * time.Second)
			continue
		}
		var addrs []resolver.Address
		for _, kv := range resp.Kvs {
			addrs = append(addrs, resolver.Address{Addr: string(kv.Value)})
		}
		r.cc.NewAddress(addrs)
		time.Sleep(5 * time.Second)
	}
}
