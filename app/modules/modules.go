package modules

import (
	"github.com/mashbens/cps/api"
	"github.com/mashbens/cps/api/v1/auth"
	"github.com/mashbens/cps/api/v1/user"
	"github.com/mashbens/cps/config"
	"github.com/mashbens/cps/util"

	authService "github.com/mashbens/cps/business/user"
	jwtService "github.com/mashbens/cps/business/user"
	userService "github.com/mashbens/cps/business/user"
	userRepo "github.com/mashbens/cps/repository/user"
)

func RegisterModules(dbCon *util.DatabaseConnection, config *config.AppConfig) api.Controller {
	userRepo := userRepo.RepositoryFactory(dbCon)

	userService := userService.NewUserService(userRepo)
	jwtService := jwtService.NewJWTService()
	authService := authService.NewAuthService(userService, jwtService)

	controller := api.Controller{
		UserAuth: auth.NewAuthController(authService, userService),
		User:     user.NewUserController(userService, jwtService),
	}
	return controller
}
