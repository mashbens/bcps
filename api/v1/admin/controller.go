package admin

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"

	"github.com/labstack/echo/v4"
	_response "github.com/mashbens/cps/api/common/response"
	"github.com/mashbens/cps/api/v1/admin/request"
	"github.com/mashbens/cps/api/v1/admin/resp"

	service "github.com/mashbens/cps/business/admin"
	jwtService "github.com/mashbens/cps/business/user"
)

type AdminController struct {
	admin      service.AdminService
	jwtService jwtService.JWTService
}

func NewAdminController(
	admin service.AdminService,
	jwtService jwtService.JWTService,

) *AdminController {
	return &AdminController{
		admin:      admin,
		jwtService: jwtService,
	}
}

func (controller *AdminController) RegisteAdmin(c echo.Context) error {
	var newAdmin request.AdminReq
	header := c.Request().Header.Get("Authorization")
	err := c.Bind(&newAdmin)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	token := controller.jwtService.ValidateToken(header, c)
	if header == "" {
		response := _response.BuildErrorResponse("Failed to process request", "Failed to validate token", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if token == nil {
		response := _response.BuildErrorResponse("Failed to process request", "Failed to validate token", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	adminintID, err := strconv.Atoi(id)

	newAdmin.SuperAdminID = adminintID

	admin, err := controller.admin.InsertAdmin(request.NewAdminReq(newAdmin))
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	data := resp.FromServiceAdmin(*admin)
	_response := _response.BuildSuccsessResponse("Super Admin created successfully", true, data)
	return c.JSON(http.StatusOK, _response)
}

func (controller *AdminController) LoginAdmin(c echo.Context) error {
	var newAdmin request.AdminReq
	err := c.Bind(&newAdmin)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	admin, err := controller.admin.AdminLogin(request.NewAdminReq(newAdmin))
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	data := resp.FromServiceAdmin(*admin)
	_response := _response.BuildSuccsessResponse("User created successfully", true, data)
	return c.JSON(http.StatusOK, _response)
}
