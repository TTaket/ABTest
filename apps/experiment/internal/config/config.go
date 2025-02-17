package config

import "flag"

var (
	Port = flag.Int("port", 50051, "The server port")
)
