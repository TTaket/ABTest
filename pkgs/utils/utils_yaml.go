package utils

import (
	"os"

	"gopkg.in/yaml.v2"
)

// ReadYamlFile reads a yaml file and unmarshals it into the provided interface
func ReadYamlFile(file string, v interface{}) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, v)
}
