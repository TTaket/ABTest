package main

import (
	dig "ABTest/apps/Micro/experiment/dig"
	conf "ABTest/apps/Micro/experiment/internal/config"
	commands "ABTest/pkgs/commands"
	"ABTest/pkgs/xmysql"
)

const serverName = conf.ServerName

func init() {
	conf := conf.GetConfig()
	xmysql.NewMysqlClient(conf.Mysql)
}

func main() {
	err := dig.Invoke(func(srv commands.MainInstance) {
		commands.Run(srv)
	})
	if err != nil {
		panic(err)
	}
}
