package member

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	_response "github.com/mashbens/cps/api/common/response"
	"github.com/mashbens/cps/api/v1/member/request"
	"github.com/mashbens/cps/api/v1/member/resp"
	service "github.com/mashbens/cps/business/member"
	jwtService "github.com/mashbens/cps/business/user"
)

type MemberController struct {
	memberService service.MemberService
	jwtService    jwtService.JWTService
}

func NewMemberController(
	memberService service.MemberService,
	jwtService jwtService.JWTService,

) *MemberController {
	return &MemberController{
		memberService: memberService,
		jwtService:    jwtService,
	}
}

func (controller *MemberController) CreateMember(c echo.Context) error {
	var newMember request.CreatMemberRequest
	header := c.Request().Header.Get("Authorization")
	err := c.Bind(&newMember)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request asd", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if newMember.Type == "" || newMember.Price == 0 || newMember.Duration == 0 || newMember.Description == "" {
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
	newMember.ImgBB = file

	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	sAdminintID, err := strconv.Atoi(id)
	newMember.Super_adminID = sAdminintID
	member, err := controller.memberService.CreateMemberships(request.NewCreateMemberReq(newMember))
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	data := resp.FromService(*member)
	_response := _response.BuildSuccsessResponse("User created successfully", true, data)
	return c.JSON(http.StatusOK, _response)
}

func (controller *MemberController) GetAllMemberType(c echo.Context) error {
	res := controller.memberService.FindAllMemberType("")

	data := resp.FromServiceSlice(res)

	_response := _response.BuildSuccsessResponse("All Membership types", true, data)
	return c.JSON(http.StatusOK, _response)
}

func (controller *MemberController) FindMemberByID(c echo.Context) error {
	id := c.Param("id")
	log.Println(id)

	member, err := controller.memberService.FindMemberTypeByID(id)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	data := resp.FromService(*member)
	_response := _response.BuildSuccsessResponse("Member found", true, data)
	return c.JSON(http.StatusOK, _response)

}
func (controller *MemberController) UpdateMemberType(c echo.Context) error {
	var newMember request.CreatMemberRequest
	header := c.Request().Header.Get("Authorization")
	err := c.Bind(&newMember)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if newMember.Type == "" || newMember.Price == 0 || newMember.Duration == 0 || newMember.Description == "" {
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
	newMember.ID = intID

	file, err := c.FormFile("image")
	if err != nil {
		return err
	}
	newMember.ImgBB = file

	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	sAdminintID, err := strconv.Atoi(id)
	newMember.Super_adminID = sAdminintID
	member, err := controller.memberService.UpdateMemberType(request.NewCreateMemberReq(newMember))
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	data := resp.FromService(*member)
	_response := _response.BuildSuccsessResponse("Member found", true, data)
	return c.JSON(http.StatusOK, _response)
}

func (controller *MemberController) DeleteMemberType(c echo.Context) error {
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

	member := controller.memberService.DeleteMemberType(adminID, memberID)
	_ = member

	_response := _response.BuildSuccsessResponse("Member Deleted", true, nil)
	return c.JSON(http.StatusOK, _response)

}
