package main

import (
	s "ABTest/apps/experiment/internal/standserver"
	commands "ABTest/pkgs/commands"
)

func main() {
	commands.Run(s.NewStandServer())
}
