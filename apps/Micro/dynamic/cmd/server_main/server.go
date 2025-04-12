package main

import (
	dig "ABTest/apps/Micro/dynamic/dig"
	conf "ABTest/apps/Micro/dynamic/internal/config"
	commands "ABTest/pkgs/commands"
)

const serverName = conf.ServerName

func init() {
}

func main() {
	err := dig.Invoke(func(srv commands.MainInstance) {
		commands.Run(srv)
	})
	if err != nil {
		panic(err)
	}
}
