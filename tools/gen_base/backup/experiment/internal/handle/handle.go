package handle

import (
	experimentServer "ABTest/apps/Micro/experiment/internal/server"
	commands "ABTest/pkgs/commands"
)

func NewHandle(s experimentServer.ExperimentServer) commands.MainInstance {
	return &handle{srv: s}
}

// handle 负责启动服务和生命周期管理
type handle struct {
	srv experimentServer.ExperimentServer
}

func (s *handle) Initialize() error {

	return nil
}
func (s *handle) RunLoop() {
	s.srv.Run()
}
func (s *handle) Destroy() {
}
