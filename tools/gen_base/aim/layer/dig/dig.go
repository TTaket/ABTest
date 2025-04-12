package dig

import (
	"ABTest/apps/Micro/layer/internal/config"
	"ABTest/apps/Micro/layer/internal/handler"
	"ABTest/apps/Micro/layer/internal/server"
	"ABTest/apps/Micro/layer/internal/service"
	"log"

	"go.uber.org/dig"
)

var container = dig.New()

func init() {
	Provide(handler.NewHandler)
	Provide(server.NewLayerServer)
	Provide(service.NewLayerService)
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
