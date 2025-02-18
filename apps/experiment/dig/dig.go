package dig

import (
	"ABTest/apps/experiment/internal/config"
	"ABTest/apps/experiment/internal/server"
	"ABTest/apps/experiment/internal/service"
	"ABTest/apps/experiment/internal/standserver"
	"log"

	"go.uber.org/dig"
)

var container = dig.New()

func init() {
	Provide(standserver.NewStandServer)
	Provide(server.NewExperimentServer)
	Provide(service.NewExperimentService)
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
