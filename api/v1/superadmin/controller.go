package superadmin

import (
	"net/http"

	"github.com/labstack/echo/v4"
	_response "github.com/mashbens/cps/api/common/response"
	"github.com/mashbens/cps/api/v1/superadmin/request"
	"github.com/mashbens/cps/api/v1/superadmin/resp"

	service "github.com/mashbens/cps/business/superadmin"
)

type SuperAdminController struct {
	superAdmin service.SuperAdminService
}

func NewSuperAdminController(
	superAdmin service.SuperAdminService,
) *SuperAdminController {
	return &SuperAdminController{
		superAdmin: superAdmin,
	}
}

func (controller *SuperAdminController) RegisterSuperAdmin(c echo.Context) error {
	var newSuperAdmin request.SuperAdminReq
	err := c.Bind(&newSuperAdmin)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if newSuperAdmin.Password == "" || newSuperAdmin.Name == "" {
		response := _response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	superAdmin, err := controller.superAdmin.CreateSuperAdmin(request.NewSuperAdminReq(newSuperAdmin))
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	data := resp.FromServiceSuperAdmin(*superAdmin)
	_response := _response.BuildSuccsessResponse("Super Admin created successfully", true, data)
	return c.JSON(http.StatusOK, _response)
}

func (controller *SuperAdminController) LoginSuperAdmin(c echo.Context) error {
	var newSuperAdmin request.SuperAdminReq
	err := c.Bind(&newSuperAdmin)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if newSuperAdmin.Password == "" || newSuperAdmin.Name == "" {
		response := _response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	superAdmin, err := controller.superAdmin.Login(request.NewSuperAdminReq(newSuperAdmin))
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	data := resp.FromServiceSuperAdmin(*superAdmin)
	_response := _response.BuildSuccsessResponse("User created successfully", true, data)
	return c.JSON(http.StatusOK, _response)
}
