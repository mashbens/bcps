package member

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	_response "github.com/mashbens/cps/api/common/response"
	"github.com/mashbens/cps/api/v1/member/request"
	"github.com/mashbens/cps/api/v1/member/resp"
	service "github.com/mashbens/cps/business/member"
	jwtService "github.com/mashbens/cps/business/user"

	"github.com/labstack/echo/v4"
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
	sAdminintID, err := strconv.Atoi(id)
	newMember.Super_adminID = sAdminintID
	member, err := controller.memberService.CreateMemberships(request.NewCreateMemberReq(newMember))
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	data := resp.FromService(*member)
	_response := _response.BuildSuccsessResponse("User created successfully", true, data)
	return c.JSON(http.StatusOK, _response)
}

func (controller *MemberController) GetAllMemberType(c echo.Context) error {
	res := controller.memberService.FIndAllMemberType("")

	data := resp.FromServiceSlice(res)

	_response := _response.BuildSuccsessResponse("All Membership types", true, data)
	return c.JSON(http.StatusOK, _response)
}
