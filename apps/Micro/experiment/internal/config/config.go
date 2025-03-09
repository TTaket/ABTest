package config

import (
	conf "ABTest/pkgs/config"
	log "ABTest/pkgs/logger"
	utils "ABTest/pkgs/utils"
	"flag"

	"github.com/spf13/cast"
)

const ServerName = "experiment"

type Config struct {
	Name string `yaml:"name"`
	Desc string `yaml:"description"`
	ID   int    `yaml:"server_id"`
	Log  string `yaml:"log"`

	Grpc *conf.Grpc `yaml:"grpc_server"`
	Etcd *conf.Etcd `yaml:"etcd"`
}

var (
	confFile = flag.String("experiment_confpos", "./configs/experiment.yml", "The configuration file path")
	grpcPort = flag.Int("experiment_port", 50051, "The server port")
	serverID = flag.Int("experiment_id", 1, "The server id")
)
var (
	config = new(Config)
	bInit  = false
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
		log.MakeLogger(config.Log, GetConfig().Name+cast.ToString(GetConfig().ID))
	}
}

func init() {
	doinit()
}

func GetConfig() *Config {
	doinit()
	return config
}
