package main

import (
	dig "ABTest/apps/experiment/dig"
	commands "ABTest/pkgs/commands"
)

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
