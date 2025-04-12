package handler

import (
	Server "ABTest/apps/Micro/dynamic/internal/server"
	commands "ABTest/pkgs/commands"
)

func NewHandler(s Server.DynamicServer) commands.MainInstance {
	return &handler{srv: s}
}

// handler 负责启动服务和生命周期管理
type handler struct {
	srv Server.DynamicServer
}

func (s *handler) Initialize() error {

	return nil
}
func (s *handler) RunLoop() {
	s.srv.Run()
}
func (s *handler) Destroy() {
}
