package classoff

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	_response "github.com/mashbens/cps/api/common/response"
	"github.com/mashbens/cps/api/v1/classoff/request"
	"github.com/mashbens/cps/api/v1/classoff/resp"
	service "github.com/mashbens/cps/business/classoff"
	jwtService "github.com/mashbens/cps/business/user"

	"github.com/labstack/echo/v4"
)

type ClassOffController struct {
	classOffService service.ClassOffService
	jwtService      jwtService.JWTService
}

func NewClassController(
	classOffService service.ClassOffService,
	jwtService jwtService.JWTService,
) *ClassOffController {
	return &ClassOffController{
		classOffService: classOffService,
		jwtService:      jwtService,
	}
}

func (controller *ClassOffController) CreateClassOffline(c echo.Context) error {
	var newClass request.CreatClassReq
	header := c.Request().Header.Get("Authorization")
	err := c.Bind(&newClass)
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
	newClass.AdminID = adminintID

	class, err := controller.classOffService.InserClassOff(request.NewCreateClassReq(newClass))
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	data := resp.FromService(*class)
	_response := _response.BuildSuccsessResponse("Class created successfully", true, data)
	return c.JSON(http.StatusOK, _response)
}

func (controller *ClassOffController) GetClassOfflineByID(c echo.Context) error {

	id := c.Param("id")

	class, err := controller.classOffService.FindClassOffByID(id)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	data := resp.FromService(*class)
	_response := _response.BuildSuccsessResponse("Member found", true, data)
	return c.JSON(http.StatusOK, _response)
}

func (controller *ClassOffController) GetAllClassOffline(c echo.Context) error {
	res := controller.classOffService.FindAllClassOff("")

	data := resp.FromServiceSlice(res)

	_response := _response.BuildSuccsessResponse("All Membership types", true, data)
	return c.JSON(http.StatusOK, _response)
}
func (controller *ClassOffController) UpdateClassOffline(c echo.Context) error {
	var newClass request.CreatClassReq
	header := c.Request().Header.Get("Authorization")
	err := c.Bind(&newClass)
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
	paramId := c.Param("id")
	intID, err := strconv.Atoi(paramId)
	newClass.ID = intID

	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	adminintID, err := strconv.Atoi(id)
	newClass.AdminID = adminintID

	class, err := controller.classOffService.UpdateClassOff(request.NewCreateClassReq(newClass))
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	data := resp.FromService(*class)
	_response := _response.BuildSuccsessResponse("Member found", true, data)
	return c.JSON(http.StatusOK, _response)
}

func (controller *ClassOffController) DeleteClassOffline(c echo.Context) error {
	header := c.Request().Header.Get("Authorization")
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
	adminID := fmt.Sprintf("%v", claims["user_id"])
	memberID := c.Param("id")

	member := controller.classOffService.DeleteClassOff(adminID, memberID)
	_ = member

	_response := _response.BuildSuccsessResponse("Class Deleted", true, nil)
	return c.JSON(http.StatusOK, _response)
}
