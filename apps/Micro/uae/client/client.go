package client

import (
	"ABTest/pkgs/xgrpc"

	"google.golang.org/grpc"

	// TODO: xxx -> servername
	conf "ABTest/apps/Micro/xxx/internal/config"
)

type XxxClient interface {
	// TODO 请根据实际情况修改函数名
}

func NewXxxClient() XxxClient {
	return &XxxClient{}
}

type XxxClient struct {
}

func getClientConn() (*grpc.ClientConn, error) {
	// 填装Client配置
	configs, iocloser := xgrpc.NewGrpcClientConfigs(conf.GetConfig().Name, conf.GetConfig().Grpc, conf.GetConfig().Etcd, conf.GetConfig().Grpc.Jaeger)
	defer iocloser.Close()

	// 获取连接并返回
	return xgrpc.GetClientConn(configs)
}
