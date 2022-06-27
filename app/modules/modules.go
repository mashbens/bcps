package modules

import (
	"github.com/mashbens/cps/api"
	"github.com/mashbens/cps/api/v1/admin"
	"github.com/mashbens/cps/api/v1/auth"
	"github.com/mashbens/cps/api/v1/member"
	"github.com/mashbens/cps/api/v1/payment"
	"github.com/mashbens/cps/api/v1/superadmin"
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

	superAdminService "github.com/mashbens/cps/business/superadmin"
	superAdminRepo "github.com/mashbens/cps/repository/superadmin"

	AdminService "github.com/mashbens/cps/business/admin"
	AdminRepo "github.com/mashbens/cps/repository/admin"
)

func RegisterModules(dbCon *util.DatabaseConnection, config *config.AppConfig) api.Controller {
	userRepo := userRepo.RepositoryFactory(dbCon)

	userService := userService.NewUserService(userRepo)
	jwtService := jwtService.NewJWTService()
	authService := authService.NewAuthService(userService, jwtService)

	superAdminRepo := superAdminRepo.SuperAdminRepositoryFactory(dbCon)
	superAdminService := superAdminService.NewSuperAdminService(superAdminRepo, jwtService)

	AdminRepo := AdminRepo.AdminRepositoryFactory(dbCon)
	AdminService := AdminService.NewAdminService(AdminRepo, jwtService, superAdminService)

	memberRepo := memberRepo.MemberRepoFactory(dbCon)
	memberService := memberService.NewMemberService(memberRepo, superAdminService)

	paymentRepo := paymentRepo.PaymentRepositoryFactory(dbCon)
	paymentService := paymentService.NewPaymentService(paymentRepo, memberService, userService)

	controller := api.Controller{
		UserAuth:   auth.NewAuthController(authService, userService),
		User:       user.NewUserController(userService, jwtService),
		Payment:    payment.NewPaymentController(paymentService, jwtService),
		SuperAdmin: superadmin.NewSuperAdminController(superAdminService),
		Member:     member.NewMemberController(memberService, jwtService),
		Admin:      admin.NewAdminController(AdminService, jwtService),
	}
	return controller
}
