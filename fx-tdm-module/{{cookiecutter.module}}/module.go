package http_example

import (
	"example/src/module/module/http-example/controller"
	"example/src/module/module/http-example/repository"
	"example/src/module/module/http-example/service"

	"github.com/sigmaott/gest/package/extension/echofx"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("example",
		fx.Provide(
			echofx.AsRoute(controller.NewExampleController),
		),
		fx.Provide(
			service.NewExampleService,
		),
		fx.Provide(
			repository.NewExampleRepository,
		),
	)
}
