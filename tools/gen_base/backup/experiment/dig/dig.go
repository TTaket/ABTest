package dig

import (
	"ABTest/apps/Micro/experiment/internal/config"
	"ABTest/apps/Micro/experiment/internal/handle"
	"ABTest/apps/Micro/experiment/internal/server"
	"ABTest/apps/Micro/experiment/internal/service"
	"log"

	"go.uber.org/dig"
)

var container = dig.New()

func init() {
	Provide(handle.NewHandle)
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
