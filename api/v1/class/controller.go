package class

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	_response "github.com/mashbens/cps/api/common/response"
	"github.com/mashbens/cps/api/v1/class/request"
	"github.com/mashbens/cps/api/v1/class/resp"
	service "github.com/mashbens/cps/business/class"
	jwtService "github.com/mashbens/cps/business/user"

	"github.com/labstack/echo/v4"
)

type ClassController struct {
	classService service.ClassService
	jwtService   jwtService.JWTService
}

func NewClassController(
	classService service.ClassService,
	jwtService jwtService.JWTService,
) *ClassController {
	return &ClassController{
		classService: classService,
		jwtService:   jwtService,
	}
}

// find all class
func (controller *ClassController) GetAllClass(c echo.Context) error {
	res := controller.classService.FindAllClass("")

	data := resp.FromServiceSlice(res)

	_response := _response.BuildSuccsessResponse("All Class found", true, data)
	return c.JSON(http.StatusOK, _response)
}

func (controller *ClassController) GetClassByID(c echo.Context) error {
	id := c.Param("id")

	class, err := controller.classService.FindClassByID(id)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	data := resp.FromService(*class)
	_response := _response.BuildSuccsessResponse("Member found", true, data)
	return c.JSON(http.StatusOK, _response)
}

// find all online class
func (controller *ClassController) GetAllClasOnline(c echo.Context) error {
	res := controller.classService.FindAllClassOn("")

	data := resp.FromServiceSlice(res)

	_response := _response.BuildSuccsessResponse("All Class found", true, data)
	return c.JSON(http.StatusOK, _response)
}

// find all offline class
func (controller *ClassController) GetAllClasOffline(c echo.Context) error {
	res := controller.classService.FindAllClassOff("")

	data := resp.FromServiceSlice(res)

	_response := _response.BuildSuccsessResponse("All Class found", true, data)
	return c.JSON(http.StatusOK, _response)
}

// find online class by id
func (controller *ClassController) GetClassOnlineByID(c echo.Context) error {

	id := c.Param("id")

	class, err := controller.classService.FindClassOnByID(id)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	data := resp.FromService(*class)
	_response := _response.BuildSuccsessResponse("Member found", true, data)
	return c.JSON(http.StatusOK, _response)
}

// find offline class by id
func (controller *ClassController) GetClassOfflineByID(c echo.Context) error {

	id := c.Param("id")

	class, err := controller.classService.FindClassOffByID(id)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	data := resp.FromService(*class)
	_response := _response.BuildSuccsessResponse("Member found", true, data)
	return c.JSON(http.StatusOK, _response)
}

// insert classs
func (controller *ClassController) CreateClass(c echo.Context) error {
	var newClass request.CreatClassReq
	header := c.Request().Header.Get("Authorization")
	err := c.Bind(&newClass)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if newClass.Classname == "" || newClass.ClasType == "" || newClass.Capacity == 0 || newClass.Date == "" {
		response := _response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if newClass.Clock == "" || newClass.Trainer == "" || newClass.Duration == 0 || newClass.Description == "" {
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

	file, err := c.FormFile("image")
	if err != nil {
		return err
	}
	newClass.ImgBB = file
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	adminintID, err := strconv.Atoi(id)
	newClass.AdminID = adminintID

	class, err := controller.classService.InsertClass(request.NewCreateClassReq(newClass))
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	data := resp.FromService(*class)
	_response := _response.BuildSuccsessResponse("Class created successfully", true, data)
	return c.JSON(http.StatusOK, _response)
}

// update classs

func (controller *ClassController) UpdateClass(c echo.Context) error {
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
		return c.JSON(http.StatusUnauthorized, response)
	}
	if token == nil {
		response := _response.BuildErrorResponse("Failed to process request", "Failed to validate token", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}
	paramId := c.Param("id")
	intID, err := strconv.Atoi(paramId)
	newClass.ID = intID

	file, err := c.FormFile("image")
	if err != nil {
		return err
	}
	newClass.ImgBB = file

	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	adminintID, err := strconv.Atoi(id)
	newClass.AdminID = adminintID

	class, err := controller.classService.UpdateClass(request.NewCreateClassReq(newClass))
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	data := resp.FromService(*class)
	_response := _response.BuildSuccsessResponse("Member found", true, data)
	return c.JSON(http.StatusOK, _response)
}

// delete classs

func (controller *ClassController) DeleteClass(c echo.Context) error {
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
	adminID := fmt.Sprintf("%v", claims["user_id"])
	memberID := c.Param("id")

	member := controller.classService.DeleteClass(adminID, memberID)
	_ = member

	_response := _response.BuildSuccsessResponse("Class Deleted", true, nil)
	return c.JSON(http.StatusOK, _response)
}
