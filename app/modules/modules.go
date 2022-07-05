package modules

import (
	"github.com/mashbens/cps/api"
	"github.com/mashbens/cps/api/v1/admin"
	"github.com/mashbens/cps/api/v1/auth"
	"github.com/mashbens/cps/api/v1/booking"
	"github.com/mashbens/cps/api/v1/class"
	"github.com/mashbens/cps/api/v1/member"
	"github.com/mashbens/cps/api/v1/newsletter"
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

	classService "github.com/mashbens/cps/business/class"
	classRepo "github.com/mashbens/cps/repository/class"

	newsletterService "github.com/mashbens/cps/business/newsletter"
	newsletterRepo "github.com/mashbens/cps/repository/newsletter"

	bookingService "github.com/mashbens/cps/business/booking"
	bookingRepo "github.com/mashbens/cps/repository/booking"
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

	classRepo := classRepo.ClassRepoFactory(dbCon)
	classService := classService.NewClassService(classRepo, adminService)

	bookingRepo := bookingRepo.BookingRepoFactory(dbCon)
	bookingService := bookingService.NewBookingService(bookingRepo, userService, classService)

	newsletterRepo := newsletterRepo.NewsRepoFactory(dbCon)
	newsletterService := newsletterService.NewNewsService(newsletterRepo, adminService)

	controller := api.Controller{
		UserAuth:   auth.NewAuthController(authService, userService),
		User:       user.NewUserController(userService, jwtService),
		Payment:    payment.NewPaymentController(paymentService, jwtService),
		SuperAdmin: superadmin.NewSuperAdminController(superAdminService),
		Member:     member.NewMemberController(memberService, jwtService),
		Admin:      admin.NewAdminController(adminService, jwtService),
		Class:      class.NewClassController(classService, jwtService),
		Booking:    booking.NewBookingController(bookingService, jwtService),
		Newsletter: newsletter.NewNewsController(newsletterService, jwtService),
	}
	return controller
}
