package xetcd

import (
	"context"
	"time"

	log "ABTest/pkgs/logger"

	clientv3 "go.etcd.io/etcd/client/v3"
	resolver "google.golang.org/grpc/resolver"
)

// build interface
type etcdBuilder struct {
	EtcdClient *clientv3.Client
	scheme     string
	servername string
}
type etcdResolver struct {
	prefix     string
	target     resolver.Target
	cc         resolver.ClientConn
	etcdClient *clientv3.Client
}

// NewEtcdBuilder returns a new etcdBuilder
func NewEtcdBuilder(etcdClient *clientv3.Client, scheme string, servername string) *etcdBuilder {
	return &etcdBuilder{EtcdClient: etcdClient, scheme: scheme, servername: servername}
}

// build function
func (b *etcdBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &etcdResolver{
		target:     target,
		cc:         cc,
		etcdClient: b.EtcdClient,
		prefix:     GetPrefix(b.scheme, b.servername),
	}
	go r.watch()
	return r, nil

}

func (b *etcdBuilder) Scheme() string {
	return b.scheme
}

// resolve function
func (r *etcdResolver) ResolveNow(o resolver.ResolveNowOptions) {
	r.watch()
}
func (r *etcdResolver) Close() {}

// TODO 并发性和回调要处理
func (r *etcdResolver) watch() {
	// watch
	for {
		// get from etcd server
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		resp, err := r.etcdClient.Get(ctx, r.prefix, clientv3.WithPrefix())
		cancel()
		if err != nil {
			log.Errorf("etcd watch error %v", err)
			r.cc.ReportError(err)
			time.Sleep(2 * time.Second)
			continue
		}

		// get address
		var addrs []resolver.Address
		for _, kv := range resp.Kvs {
			addrs = append(addrs, resolver.Address{Addr: string(kv.Value)})
		}

		// update address
		r.cc.NewAddress(addrs)
		time.Sleep(5 * time.Second)
	}
}
