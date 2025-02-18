package config

import (
	conf "ABTest/pkgs/config"
	utils "ABTest/pkgs/utils"
	"flag"
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
)
var (
	config = new(Config)
)

func init() {
	flag.Parse()
	utils.ReadYamlFile(*confFile, &config)
	config.Grpc.Port = *grpcPort
}

func GetConfig() *Config {
	return config
}
