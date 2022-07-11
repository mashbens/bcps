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
	adminService service.AdminService
	jwtService   jwtService.JWTService
}

func NewAdminController(
	adminService service.AdminService,
	jwtService jwtService.JWTService,

) *AdminController {
	return &AdminController{
		adminService: adminService,
		jwtService:   jwtService,
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

	if newAdmin.Email == "" || newAdmin.Password == "" || newAdmin.Name == "" || newAdmin.Phone == "" {
		response := _response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	token := controller.jwtService.ValidateToken(header, c)
	if header == "" {
		response := _response.BuildErrorResponse("Failed to process request", "Failed to validate token", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}
	if token == nil {
		response := _response.BuildErrorResponse("Failed to process request", "Failed to validate token", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}

	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	adminintID, err := strconv.Atoi(id)

	newAdmin.SuperAdminID = adminintID

	admin, err := controller.adminService.InsertAdmin(request.NewAdminReq(newAdmin))
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	data := resp.FromServiceAdmin(*admin)
	_response := _response.BuildSuccsessResponse("Admin created successfully", true, data)
	return c.JSON(http.StatusOK, _response)
}

func (controller *AdminController) LoginAdmin(c echo.Context) error {
	var newAdmin request.AdminReq
	err := c.Bind(&newAdmin)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if newAdmin.Email == "" || newAdmin.Password == "" {
		response := _response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	admin, err := controller.adminService.AdminLogin(request.NewAdminReq(newAdmin))
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	data := resp.FromServiceAdmin(*admin)
	_response := _response.BuildSuccsessResponse("Admin login successfully", true, data)
	return c.JSON(http.StatusOK, _response)
}

func (controller *AdminController) FindAdminByID(c echo.Context) error {
	header := c.Request().Header.Get("Authorization")
	token := controller.jwtService.ValidateToken(header, c)
	if header == "" {
		response := _response.BuildErrorResponse("Failed to process request", "Failed to validate token", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}
	if token == nil {
		response := _response.BuildErrorResponse("Failed to process request", "Failed to validate token", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}
	claims := token.Claims.(jwt.MapClaims)
	_ = fmt.Sprintf("%v", claims["user_id"])
	adminID := c.Param("id")

	admin, err := controller.adminService.FindAdminByID(adminID)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	data := resp.FromServiceAdmin(*admin)
	_response := _response.BuildSuccsessResponse("Admin found", true, data)
	return c.JSON(http.StatusOK, _response)
}

func (controller *AdminController) FindAllAdmins(c echo.Context) error {
	header := c.Request().Header.Get("Authorization")
	token := controller.jwtService.ValidateToken(header, c)
	if header == "" {
		response := _response.BuildErrorResponse("Failed to process request", "Failed to validate token", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}
	if token == nil {
		response := _response.BuildErrorResponse("Failed to process request", "Failed to validate token", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}
	claims := token.Claims.(jwt.MapClaims)
	sAdminID := fmt.Sprintf("%v", claims["user_id"])
	adminID := c.Param("id")

	res := controller.adminService.FindAllAdmins(sAdminID, adminID)
	data := resp.FromServiceSlice(res)

	_response := _response.BuildSuccsessResponse("All admins found", true, data)
	return c.JSON(http.StatusOK, _response)
}
func (controller *AdminController) UpdateAdmin(c echo.Context) error {
	var newAdmin request.AdminReq
	header := c.Request().Header.Get("Authorization")
	err := c.Bind(&newAdmin)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if newAdmin.Email == "" || newAdmin.Password == "" || newAdmin.Name == "" || newAdmin.Phone == "" {
		response := _response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	token := controller.jwtService.ValidateToken(header, c)
	if header == "" {
		response := _response.BuildErrorResponse("Failed to process request", "Failed to validate token", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}
	if token == nil {
		response := _response.BuildErrorResponse("Failed to process request", "Failed to validate token", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}

	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	sAdminintID, err := strconv.Atoi(id)
	newAdmin.SuperAdminID = sAdminintID

	paramID := c.Param("id")
	reqParam, err := strconv.Atoi(paramID)
	newAdmin.ID = reqParam
	admin, err := controller.adminService.UpdateAdmin(request.NewAdminReq(newAdmin))
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	data := resp.FromServiceAdmin(*admin)
	_response := _response.BuildSuccsessResponse("Admin Updated successfully", true, data)
	return c.JSON(http.StatusOK, _response)
}
func (controller *AdminController) DeleteAdmin(c echo.Context) error {
	header := c.Request().Header.Get("Authorization")
	token := controller.jwtService.ValidateToken(header, c)
	if header == "" {
		response := _response.BuildErrorResponse("Failed to process request", "Failed to validate token", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}
	if token == nil {
		response := _response.BuildErrorResponse("Failed to process request", "Failed to validate token", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}
	claims := token.Claims.(jwt.MapClaims)
	sAdminID := fmt.Sprintf("%v", claims["user_id"])
	adminID := c.Param("id")

	member := controller.adminService.DeleteAdmin(sAdminID, adminID)
	if member != nil {
		response := _response.BuildErrorResponse("Failed to process request", "Admin not found", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	_response := _response.BuildSuccsessResponse("Admin Deleted", true, nil)
	return c.JSON(http.StatusOK, _response)
}
