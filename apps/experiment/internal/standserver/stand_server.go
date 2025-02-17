package standserver

import (
	experimentServer "ABTest/apps/experiment/internal/server"
	commands "ABTest/pkgs/commands"
)

func NewStandServer(s experimentServer.ExperimentServer) commands.MainInstance {
	return &standServer{srv: s}
}

// standServer 负责启动服务和生命周期管理
type standServer struct {
	srv experimentServer.ExperimentServer
}

func (s *standServer) Initialize() error {
	return nil
}
func (s *standServer) RunLoop() {
	s.srv.Run()
}
func (s *standServer) Destroy() {
}
