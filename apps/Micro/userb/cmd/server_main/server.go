package main

import (
	dig "ABTest/apps/Micro/userb/dig"
	conf "ABTest/apps/Micro/userb/internal/config"
	commands "ABTest/pkgs/commands"
	"ABTest/pkgs/xmysql"
)

const serverName = conf.ServerName

func init() {
	conf := conf.GetConfig()
	xmysql.NewMysqlClient(conf.MySql)
}

func main() {
	err := dig.Invoke(func(srv commands.MainInstance) {
		commands.Run(srv)
	})
	if err != nil {
		panic(err)
	}
}
