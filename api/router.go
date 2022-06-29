package api

import (
	"github.com/mashbens/cps/api/v1/admin"
	"github.com/mashbens/cps/api/v1/auth"
	"github.com/mashbens/cps/api/v1/class"
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
	Class      *class.ClassController
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

	memberRoutes := e.Group("/api/v1/membership")
	memberRoutes.POST("/create-member", controller.Member.CreateMember)
	memberRoutes.GET("/list", controller.Member.GetAllMemberType)
	memberRoutes.GET("/:id", controller.Member.FindMemberByID)
	memberRoutes.PUT("/:id", controller.Member.UpdateMemberType)
	memberRoutes.DELETE("/:id", controller.Member.DeleteMemberType)

	adminRoutes := e.Group("/api/v1/admin")
	adminRoutes.POST("/create-admin", controller.Admin.RegisteAdmin)
	adminRoutes.POST("/login", controller.Admin.LoginAdmin)
	adminRoutes.GET("/list", controller.Admin.FindAllAdmins)
	adminRoutes.GET("/:id", controller.Admin.FindAdminByID)
	adminRoutes.PUT("/:id", controller.Admin.UpdateAdmin)
	adminRoutes.DELETE("/:id", controller.Admin.DeleteAdmin)

	adminClass := e.Group("/api/v1/class")
	adminClass.POST("/create-class", controller.Class.CreateClass)
	adminClass.GET("/list", controller.Class.GetAllClass)
	adminClass.GET("/online/list", controller.Class.GetAllClasOnline)
	adminClass.GET("/offline/list", controller.Class.GetAllClasOffline)

	adminClass.GET("/online/:id", controller.Class.GetClassOnlineByID)
	adminClass.GET("/offline/:id", controller.Class.GetClassOfflineByID)
	adminClass.PUT("/:id", controller.Class.UpdateClass)
	adminClass.DELETE("/:id", controller.Class.DeleteClass)

}
