package xetcd

import (
	conf "ABTest/pkgs/config"
	log "ABTest/pkgs/logger"
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func GetEtcdClient(etcd *conf.Etcd) (*clientv3.Client, error) {
	if etcd == nil {
		log.Errorf("etcd config is nil")
		return nil, errors.New("etcd config is nil")
	}
	// etcd client
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   etcd.Endpoints,
		DialTimeout: time.Duration(etcd.DialTimeout),
	})
	if err != nil {
		log.Errorf("etcd client error: %v", err)
		return nil, err
	}
	return cli, err
}

const (
	TIME_TO_LIVE = 10 // TTL(生存时间值)
)

type RegEtcd struct {
	cli          *clientv3.Client
	endpoints    []string
	serviceValue string
	serviceKey   string
	ttl          int
	host         string
	port         int
	schema       string
	serviceName  string
	reChan       chan struct{}
	connected    bool
}

var (
	rEtcd *RegEtcd
)

// "%s:///%s/"
func GetPrefix(schema, serviceName string) string {
	return fmt.Sprintf("%s:///%s/", schema, serviceName)
}

func RegisterEtcd(schema string, endpoints []string, myHost string, myPort int, serviceName string, ttl int) (err error) {
	serviceValue := net.JoinHostPort(myHost, strconv.Itoa(myPort))
	serviceKey := GetPrefix(schema, serviceName) + serviceValue
	rEtcd = &RegEtcd{
		endpoints:    endpoints,
		serviceValue: serviceValue,
		serviceKey:   serviceKey,
		ttl:          ttl,
		host:         myHost,
		port:         myPort,
		schema:       schema,
		serviceName:  serviceName,
		reChan:       make(chan struct{}),
	}
	rEtcd.ReRegister()
	rEtcd.Register()
	return
}

func (r *RegEtcd) Register() (err error) {
	defer func() {
		if err != nil {
			r.reChan <- struct{}{}
		}
	}()

	var (
		cli    *clientv3.Client
		ctx    context.Context
		cancel context.CancelFunc
		resp   *clientv3.LeaseGrantResponse
		kresp  <-chan *clientv3.LeaseKeepAliveResponse
	)

	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   r.endpoints,
		DialTimeout: 5 * time.Second})
	if err != nil {
		log.Error(err.Error())
		return
	}

	ctx, cancel = context.WithCancel(context.Background())
	resp, err = cli.Grant(ctx, int64(r.ttl))
	if err != nil {
		log.Error(err.Error())
		return
	}

	if _, err = cli.Put(ctx, r.serviceKey, r.serviceValue, clientv3.WithLease(resp.ID)); err != nil {
		log.Error(err.Error())
		return
	}
	kresp, err = cli.KeepAlive(ctx, resp.ID)
	if err != nil {
		log.Error(err.Error())
		return
	}
	rEtcd.cli = cli

	go func() {
		for {
			select {
			case _, ok := <-kresp:
				r.connected = ok
				if ok {
					//log.Info("续约成功")
				} else {
					log.Error("租约失效")
					cancel()
					r.reChan <- struct{}{}
					return
				}
			}
		}
	}()
	return
}

func (r *RegEtcd) ReRegister() {
	go func() {
		var (
			ok bool
		)
		for {
			select {
			case _, ok = <-r.reChan:
				if ok == false {
					return
				}
				time.Sleep(1 * time.Second)
				r.Register()
			}
		}
	}()
}
