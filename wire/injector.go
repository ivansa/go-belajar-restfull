//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"go-belajar-restfull/config"
	"go-belajar-restfull/controller"
	"go-belajar-restfull/middleware"
	"go-belajar-restfull/repository"
	"go-belajar-restfull/service"
	"net/http"
)

var EmployeeSet = wire.NewSet(
	repository.NewEmployeeRepository,
	wire.Bind(new(repository.EmployeeRepository), new(*repository.EmployeeRepositoryImpl)),
	service.NewEmployeeService,
	wire.Bind(new(service.EmployeeService), new(*service.EmployeeServiceImpl)),
	controller.NewEmployeeController,
	wire.Bind(new(controller.EmployeeController), new(*controller.EmployeeControllerImpl)),
)

func ProvideValidator() *validator.Validate {
	return validator.New()
}

func InitializeServer() *http.Server {
	wire.Build(
		config.InitDB,
		ProvideValidator,
		EmployeeSet,
		config.NewHttpRouter,
		middleware.NewApiMiddleware,
		config.NewHttpServer,
	)
	return nil
}
