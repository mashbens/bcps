package modules

import (
	"github.com/mashbens/cps/api"
	"github.com/mashbens/cps/api/v1/auth"
	"github.com/mashbens/cps/api/v1/payment"
	"github.com/mashbens/cps/api/v1/user"
	"github.com/mashbens/cps/config"
	"github.com/mashbens/cps/util"

	authService "github.com/mashbens/cps/business/user"
	jwtService "github.com/mashbens/cps/business/user"

	userService "github.com/mashbens/cps/business/user"
	userRepo "github.com/mashbens/cps/repository/user"

	memberService "github.com/mashbens/cps/business/member"
	memberRepo "github.com/mashbens/cps/repository/member"

	paymentService "github.com/mashbens/cps/business/payment"
	paymentRepo "github.com/mashbens/cps/repository/payment"
)

func RegisterModules(dbCon *util.DatabaseConnection, config *config.AppConfig) api.Controller {
	userRepo := userRepo.RepositoryFactory(dbCon)

	userService := userService.NewUserService(userRepo)
	jwtService := jwtService.NewJWTService()
	authService := authService.NewAuthService(userService, jwtService)

	memberRepo := memberRepo.MemberRepoFactory(dbCon)
	memberService := memberService.NewMemberService(memberRepo)

	paymentRepo := paymentRepo.PaymentRepositoryFactory(dbCon)
	paymentService := paymentService.NewPaymentService(paymentRepo, memberService, userService)

	controller := api.Controller{
		UserAuth: auth.NewAuthController(authService, userService),
		User:     user.NewUserController(userService, jwtService),
		Payment:  payment.NewPaymentController(paymentService, jwtService),
	}
	return controller
}
