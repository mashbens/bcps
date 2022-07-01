package booking

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	_response "github.com/mashbens/cps/api/common/response"
	"github.com/mashbens/cps/api/v1/booking/request"
	"github.com/mashbens/cps/api/v1/booking/resp"
	service "github.com/mashbens/cps/business/booking"
	jwtService "github.com/mashbens/cps/business/user"
)

type BookingController struct {
	bookingService service.BookingService
	jwtService     jwtService.JWTService
}

func NewBookingController(
	bookingService service.BookingService,
	jwtService jwtService.JWTService,
) *BookingController {
	return &BookingController{
		bookingService: bookingService,
		jwtService:     jwtService,
	}
}

func (controller *BookingController) CreateBooking(c echo.Context) error {
	var createBookingReq request.CreateBookingRequest
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
	if err := c.Bind(&createBookingReq); err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	id := fmt.Sprintf("%v", claims["user_id"])
	intID, err := strconv.Atoi(id)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	createBookingReq.UserID = intID
	log.Println(createBookingReq.ClassID, "<-----")
	Bookingt, err := controller.bookingService.InsertBooking(request.NewCreateBookingReq(createBookingReq))
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	data := resp.FromService(*Bookingt)
	response := _response.BuildSuccsessResponse("User Booke Succsessfully", true, data)
	return c.JSON(http.StatusOK, response)
}
