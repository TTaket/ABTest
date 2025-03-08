package handle

import (
	Xxx "ABTest/apps/Xxx/internal/server"
	commands "ABTest/pkgs/commands"
)

func NewHandle(s Xxx.XxxServer) commands.MainInstance {
	return &handle{srv: s}
}

// handle 负责启动服务和生命周期管理
type handle struct {
	srv Xxx.XxxServer
}

func (s *handle) Initialize() error {

	return nil
}
func (s *handle) RunLoop() {
	s.srv.Run()
}
func (s *handle) Destroy() {
}
