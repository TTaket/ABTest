package utils

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Base struct {
	ModuleName       string `yaml:"modulename"`
	ModuleNameBig    string `yaml:"modulename_big"`
	ModuleNameSmall  string `yaml:"modulename_sma"`
	RPCServiceName   string `yaml:"rpcservicename"`
	RPCServiceClient string `yaml:"rpcserviceclient"`
	Port             string `yaml:"port"`
}

type Function struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

type Config struct {
	BaseInfo Base       `yaml:"base"`
	FuncInfo []Function `yaml:"func"`
}

// ReadYamlFile reads a yaml file and unmarshals it into the provided interface
func ReadYamlFile(file string, v interface{}) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, v)
}
