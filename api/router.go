package api

import (
	"github.com/mashbens/cps/api/v1/admin"
	"github.com/mashbens/cps/api/v1/auth"
	"github.com/mashbens/cps/api/v1/member"
	"github.com/mashbens/cps/api/v1/payment"
	"github.com/mashbens/cps/api/v1/superadmin"
	"github.com/mashbens/cps/api/v1/user"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	UserAuth   *auth.AuthController
	User       *user.UserController
	Payment    *payment.PaymentController
	SuperAdmin *superadmin.SuperAdminController
	Member     *member.MemberController
	Admin      *admin.AdminController
}

func RegisterRoutes(e *echo.Echo, controller *Controller) {
	userAuthRoutes := e.Group("/api/v1/auth")
	userAuthRoutes.POST("/register", controller.UserAuth.RegisterHandler)
	userAuthRoutes.POST("/login", controller.UserAuth.LoginHandler)
	userAuthRoutes.POST("/verify-email", controller.UserAuth.EmailVerificationHandler)
	userAuthRoutes.POST("/reset-password", controller.UserAuth.ForgotPasswordHandler)
	userAuthRoutes.POST("/reset-password-confirm", controller.UserAuth.FindUserByEmailHandler)

	userRoutes := e.Group("/api/v1/user")
	userRoutes.GET("/profile", controller.User.Profile)
	userRoutes.PUT("/profile", controller.User.Update)

	paymentRoutes := e.Group("/api/v1/member")
	paymentRoutes.POST("/register", controller.Payment.CreatePayment)
	paymentRoutes.GET("/details", controller.Payment.GetPaymentDetail)

	superAdminRoutes := e.Group("/api/v1/super-admin")
	superAdminRoutes.POST("/register", controller.SuperAdmin.RegisterSuperAdmin)
	superAdminRoutes.POST("/login", controller.SuperAdmin.LoginSuperAdmin)

	memberRoutes := e.Group("/api/v1/member")
	memberRoutes.POST("/create-member", controller.Member.CreateMember)
	memberRoutes.GET("/list", controller.Member.GetAllMemberType)

	adminRoutes := e.Group("/api/v1/admin")
	adminRoutes.POST("/create-admin", controller.Admin.RegisteAdmin)
	adminRoutes.POST("/login", controller.Admin.LoginAdmin)
}
