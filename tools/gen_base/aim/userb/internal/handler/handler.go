package handler

import (
	Server "ABTest/apps/Micro/userb/internal/server"
	commands "ABTest/pkgs/commands"
)

func NewHandler(s Server.UserbServer) commands.MainInstance {
	return &handler{srv: s}
}

// handler 负责启动服务和生命周期管理
type handler struct {
	srv Server.UserbServer
}

func (s *handler) Initialize() error {

	return nil
}
func (s *handler) RunLoop() {
	s.srv.Run()
}
func (s *handler) Destroy() {
}
