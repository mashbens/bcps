package classon

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	_response "github.com/mashbens/cps/api/common/response"
	"github.com/mashbens/cps/api/v1/classon/request"
	"github.com/mashbens/cps/api/v1/classon/resp"
	service "github.com/mashbens/cps/business/classon"
	jwtService "github.com/mashbens/cps/business/user"

	"github.com/labstack/echo/v4"
)

type ClassOnController struct {
	classOnService service.ClassOnService
	jwtService     jwtService.JWTService
}

func NewClassController(
	classOnService service.ClassOnService,
	jwtService jwtService.JWTService,
) *ClassOnController {
	return &ClassOnController{
		classOnService: classOnService,
		jwtService:     jwtService,
	}
}

func (controller *ClassOnController) CreateClassOnline(c echo.Context) error {
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

	class, err := controller.classOnService.InserClassOn(request.NewCreateClassReq(newClass))
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	data := resp.FromService(*class)
	_response := _response.BuildSuccsessResponse("Class created successfully", true, data)
	return c.JSON(http.StatusOK, _response)
}

func (controller *ClassOnController) GetClassOnlineByID(c echo.Context) error {

	id := c.Param("id")

	class, err := controller.classOnService.FindClassOnByID(id)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	data := resp.FromService(*class)
	_response := _response.BuildSuccsessResponse("Class found", true, data)
	return c.JSON(http.StatusOK, _response)
}

func (controller *ClassOnController) GetAllClasOnline(c echo.Context) error {
	res := controller.classOnService.FindAllClassOn("")

	data := resp.FromServiceSlice(res)

	_response := _response.BuildSuccsessResponse("All Class found", true, data)
	return c.JSON(http.StatusOK, _response)
}
func (controller *ClassOnController) UpdateClassOnline(c echo.Context) error {
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

	class, err := controller.classOnService.UpdateClassOn(request.NewCreateClassReq(newClass))
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	data := resp.FromService(*class)
	_response := _response.BuildSuccsessResponse("Class Updated", true, data)
	return c.JSON(http.StatusOK, _response)
}

func (controller *ClassOnController) DeleteClassOnline(c echo.Context) error {
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
	classID := c.Param("id")

	class := controller.classOnService.DeleteClassOn(adminID, classID)
	_ = class

	_response := _response.BuildSuccsessResponse("Class Deleted", true, nil)
	return c.JSON(http.StatusOK, _response)
}
