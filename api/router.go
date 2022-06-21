package api

import (
	"github.com/mashbens/cps/api/v1/auth"
	"github.com/mashbens/cps/api/v1/user"

	service "github.com/mashbens/cps/business/user"

	"github.com/labstack/echo/v4"
)

var jwtService service.JWTService = service.NewJWTService()

type Controller struct {
	UserAuth *auth.AuthController
	User     *user.UserController
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
}
