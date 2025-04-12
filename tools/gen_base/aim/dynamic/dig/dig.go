package dig

import (
	"ABTest/apps/Micro/dynamic/internal/config"
	"ABTest/apps/Micro/dynamic/internal/handler"
	"ABTest/apps/Micro/dynamic/internal/server"
	"ABTest/apps/Micro/dynamic/internal/service"
	"log"

	"go.uber.org/dig"
)

var container = dig.New()

func init() {
	Provide(handler.NewHandler)
	Provide(server.NewDynamicServer)
	Provide(service.NewDynamicService)
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
