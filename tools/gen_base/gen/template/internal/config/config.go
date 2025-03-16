package config

import (
	conf "ABTest/pkgs/config"
	log "ABTest/pkgs/logger"
	utils "ABTest/pkgs/utils"
	"flag"

	"github.com/spf13/cast"
)

const ServerName = "%1"

type Config struct {
	Name string `yaml:"name"`
	Desc string `yaml:"description"`
	ID   int    `yaml:"server_id"`
	Log  string `yaml:"log"`

	Grpc *conf.Grpc `yaml:"grpc_server"`
	Etcd *conf.Etcd `yaml:"etcd"`
}

var (
	confFile = flag.String("%3_confpos", "./configs/%1.yml", "The configuration file path")
	grpcPort = flag.Int("%3_port", 50051, "The server port")
	serverID = flag.Int("%3_id", 1, "The server id")
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
		if config.Grpc == nil {
			panic("config.Grpc is nil")
		}
		config.Grpc.Port = *grpcPort
		config.ID = *serverID
		config.Grpc.ServerID = *serverID
		Log = log.NewMyLogger(config.Log, GetConfig().Name+cast.ToString(GetConfig().ID))
	}
}

func init() {
	doinit()
}

func GetConfig() *Config {
	doinit()
	return config
}
