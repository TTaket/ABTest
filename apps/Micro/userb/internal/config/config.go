package config

import (
	conf "ABTest/pkgs/config"
	"ABTest/pkgs/logger"
	log "ABTest/pkgs/logger"
	utils "ABTest/pkgs/utils"
	"flag"

	"github.com/spf13/cast"
)

const ServerName = "userb"

// 服务内所用配置 从服务内config文件中读取
type Config struct {
	Name string `yaml:"name"`
	Desc string `yaml:"description"`
	ID   int    `yaml:"server_id"`
	Log  string `yaml:"log"`

	Grpc  *conf.Grpc  `yaml:"grpc_server"`
	Etcd  *conf.Etcd  `yaml:"etcd"`
	MySql *conf.Mysql `yaml:"mysql"`
}

// 从配置文件中读取配置并且支持命令行参数
var (
	confFile = flag.String("userb_confpos", "./configs/userb.yml", "The configuration file path")
	grpcPort = flag.Int("userb_port", 50052, "The server port")
	serverID = flag.Int("userb_id", 1, "The server id")
)
var (
	config = new(Config)
	bInit  = false
	Log    = &log.MyLogger{}
)

func doinit() {
	if bInit == false {
		bInit = true
		flag.Parse()
		utils.ReadYamlFile(*confFile, &config)

		// 根据实际情况修改
		if config.Grpc == nil {
			panic("config.Grpc is nil")
		}
		config.Grpc.Port = *grpcPort
		config.ID = *serverID
		config.Grpc.ServerID = *serverID

		// 初始化日志
		Log = logger.NewMyLogger(config.Log, GetConfig().Name+cast.ToString(GetConfig().ID))
	}
}

func init() {
	doinit()
}

func GetConfig() *Config {
	doinit()
	return config
}
