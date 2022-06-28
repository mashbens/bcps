package modules

import (
	"github.com/mashbens/cps/api"
	"github.com/mashbens/cps/api/v1/admin"
	"github.com/mashbens/cps/api/v1/auth"
	"github.com/mashbens/cps/api/v1/classoff"

	"github.com/mashbens/cps/api/v1/classon"
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

	adminService "github.com/mashbens/cps/business/admin"
	adminRepo "github.com/mashbens/cps/repository/admin"

	classOnService "github.com/mashbens/cps/business/classon"
	classOnRepo "github.com/mashbens/cps/repository/classon"

	classOffService "github.com/mashbens/cps/business/classoff"
	classOffRepo "github.com/mashbens/cps/repository/classoff"
)

func RegisterModules(dbCon *util.DatabaseConnection, config *config.AppConfig) api.Controller {
	userRepo := userRepo.RepositoryFactory(dbCon)

	userService := userService.NewUserService(userRepo)
	jwtService := jwtService.NewJWTService()
	authService := authService.NewAuthService(userService, jwtService)

	superAdminRepo := superAdminRepo.SuperAdminRepositoryFactory(dbCon)
	superAdminService := superAdminService.NewSuperAdminService(superAdminRepo, jwtService)

	adminRepo := adminRepo.AdminRepositoryFactory(dbCon)
	adminService := adminService.NewAdminService(adminRepo, jwtService, superAdminService)

	memberRepo := memberRepo.MemberRepoFactory(dbCon)
	memberService := memberService.NewMemberService(memberRepo, superAdminService)

	paymentRepo := paymentRepo.PaymentRepositoryFactory(dbCon)
	paymentService := paymentService.NewPaymentService(paymentRepo, memberService, userService)

	classOnRepo := classOnRepo.ClassOnRepoFactory(dbCon)
	classOnService := classOnService.NewClassOnService(classOnRepo, adminService)

	classOffRepo := classOffRepo.ClassOffRepoFactory(dbCon)
	classOffService := classOffService.NewClassOffService(classOffRepo, adminService)

	controller := api.Controller{
		UserAuth:   auth.NewAuthController(authService, userService),
		User:       user.NewUserController(userService, jwtService),
		Payment:    payment.NewPaymentController(paymentService, jwtService),
		SuperAdmin: superadmin.NewSuperAdminController(superAdminService),
		Member:     member.NewMemberController(memberService, jwtService),
		Admin:      admin.NewAdminController(adminService, jwtService),
		ClassOn:    classon.NewClassController(classOnService, jwtService),
		ClassOff:   classoff.NewClassController(classOffService, jwtService),
	}
	return controller
}
