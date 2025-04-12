package dig

import (
	"ABTest/apps/Micro/%1/internal/config"
	"ABTest/apps/Micro/%1/internal/handler"
	"ABTest/apps/Micro/%1/internal/server"
	"ABTest/apps/Micro/%1/internal/service"
	"log"

	"go.uber.org/dig"
)

var container = dig.New()

func init() {
	Provide(handler.NewHandler)
	Provide(server.New%2Server)
	Provide(service.New%2Service)
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
