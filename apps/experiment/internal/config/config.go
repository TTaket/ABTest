package config

import (
	conf "ABTest/pkgs/config"
	log "ABTest/pkgs/logger"
	utils "ABTest/pkgs/utils"
	"flag"

	"github.com/spf13/cast"
)

type Config struct {
	Name string `yaml:"name"`
	Desc string `yaml:"description"`
	ID   int    `yaml:"server_id"`
	Log  string `yaml:"log"`

	Grpc conf.Grpc `yaml:"grpc_server"`
}

var (
	confFile = flag.String("confpos", "./configs/experiment.yml", "The configuration file path")
	grpcPort = flag.Int("port", 50051, "The server port")
	serverID = flag.Int("id", 1, "The server id")
)
var (
	config = new(Config)
)

func init() {
	flag.Parse()
	utils.ReadYamlFile(*confFile, &config)
	config.Grpc.Port = *grpcPort
	config.ID = *serverID
	config.Grpc.ServerID = *serverID

	log.MakeLogger(config.Log, GetConfig().Name+cast.ToString(GetConfig().ID))
}

func GetConfig() *Config {
	return config
}
