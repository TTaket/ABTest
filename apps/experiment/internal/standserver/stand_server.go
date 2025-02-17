package standserver

import (
	experimentServer "ABTest/apps/experiment/internal/server"
	commands "ABTest/pkgs/commands"
)

func NewStandServer() commands.MainInstance {
	return &standServer{experimentServer.NewExperimentServer()}
}

// standServer 负责启动服务和生命周期管理
type standServer struct {
	experimentServer.ExperimentServer
}

func (s *standServer) Initialize() error {
	return nil
}
func (s *standServer) RunLoop() {
	s.Run()
}
func (s *standServer) Destroy() {
}
