package auth

import (
	"net/http"

	_response "github.com/mashbens/cps/api/common/response"
	"github.com/mashbens/cps/api/v1/auth/request"
	"github.com/mashbens/cps/api/v1/auth/resp"
	service "github.com/mashbens/cps/business/user"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authService service.AuthService
	userService service.UserService
}

func NewAuthController(
	authService service.AuthService,
	userService service.UserService,
) *AuthController {
	return &AuthController{
		authService: authService,
		userService: userService,
	}
}

func (controller *AuthController) RegisterHandler(c echo.Context) error {
	var newUser request.RegisterRequest
	err := c.Bind(&newUser)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if newUser.Email == "" || newUser.Password == "" || newUser.Name == "" || newUser.Phone == "" {
		response := _response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if len(newUser.Password) < 6 {
		response := _response.BuildErrorResponse("Failed to process request", "Password must be at least 6 characters", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if len(newUser.Phone) < 1 {
		response := _response.BuildErrorResponse("Failed to process request", "Phone must be at least 1 character", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	user, err := controller.authService.Register(request.NewRegisterRequest(newUser))
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := resp.FromService(*user)
	_response := _response.BuildSuccsessResponse("User created successfully", true, data)
	return c.JSON(http.StatusOK, _response)
}

func (controller *AuthController) LoginHandler(c echo.Context) error {
	var loginRequest request.LoginRequest
	err := c.Bind(&loginRequest)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)

	}

	if loginRequest.Email == "" || loginRequest.Password == "" {
		response := _response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	user, err := controller.authService.Login(request.NewLoginRequest(loginRequest))
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := resp.FromService(*user)
	_response := _response.BuildSuccsessResponse("User logged in successfully", true, data)
	return c.JSON(http.StatusOK, _response)
}

// ---
func (controller *AuthController) FindUserByEmailHandler(c echo.Context) error {
	var FindByEmail request.FindUserByEmailRequest
	err := c.Bind(&FindByEmail)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	user, err := controller.userService.FindUserByEmail(FindByEmail.Email)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := resp.FromService(*user)
	_response := _response.BuildSuccsessResponse("User found successfully", true, data)
	return c.JSON(http.StatusOK, _response)
}

func (controller *AuthController) EmailVerificationHandler(c echo.Context) error {
	var emailVerificationRequest request.EmailVerificationRequest
	err := c.Bind(&emailVerificationRequest)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	user, err := controller.authService.SendEmailVerification(request.NewEmailVerificationRequest(emailVerificationRequest))
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := resp.FromService(*user)
	_response := _response.BuildSuccsessResponse("User verified successfully", true, data)
	return c.JSON(http.StatusOK, _response)
}

func (controller *AuthController) ForgotPasswordHandler(c echo.Context) error {
	var forgotPasswordRequest request.PasswordResetRequest
	err := c.Bind(&forgotPasswordRequest)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	user, err := controller.authService.ResetPassword(request.NewPasswordResetRequest(forgotPasswordRequest))
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := resp.FromService(*user)
	_response := _response.BuildSuccsessResponse("User Password reset successfully", true, data)
	return c.JSON(http.StatusOK, _response)
}
