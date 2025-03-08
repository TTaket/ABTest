package dig

import (
	"ABTest/apps/Micro/userb/internal/config"
	"ABTest/apps/Micro/userb/internal/handle"
	"ABTest/apps/Micro/userb/internal/server"
	"ABTest/apps/Micro/userb/internal/service"
	"log"

	"go.uber.org/dig"
)

var container = dig.New()

func init() {
	Provide(handle.NewHandle)
	Provide(server.NewUserbServer)
	Provide(service.NewUserbService)
	Provide(config.GetConfig)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}

func Provide(constructor interface{}, opts ...dig.ProvideOption) {
	err := container.Provide(constructor)
	if err != nil {
		log.Panic(err)
	}
}
