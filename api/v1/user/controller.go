package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"

	service "github.com/mashbens/cps/business/user"

	"github.com/labstack/echo/v4"
	_response "github.com/mashbens/cps/api/common/response"
	"github.com/mashbens/cps/api/v1/user/request"
	"github.com/mashbens/cps/api/v1/user/resp"
)

type UserController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(
	userService service.UserService,
	jwtService service.JWTService,
) *UserController {
	return &UserController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (controller *UserController) Profile(c echo.Context) error {
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
	id := fmt.Sprintf("%v", claims["user_id"])
	user, err := controller.userService.FindUserByID(id)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	data := resp.FromServiceUser(*user)
	response := _response.BuildSuccsessResponse("User found", true, data)
	return c.JSON(http.StatusOK, response)
}

func (controller *UserController) Update(c echo.Context) error {
	var userReq request.UpdateUserRequest
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
	if err := c.Bind(&userReq); err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	id := fmt.Sprintf("%v", claims["user_id"])
	intID, err := strconv.Atoi(id)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	userReq.ID = intID
	user, err := controller.userService.UpdateUser(request.NewUpdateUserRequest(userReq))
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	data := resp.FromServiceUser(*user)
	response := _response.BuildSuccsessResponse("User updated", true, data)
	return c.JSON(http.StatusOK, response)

}
