package payment

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	_response "github.com/mashbens/cps/api/common/response"
	"github.com/mashbens/cps/api/v1/payment/request"
	"github.com/mashbens/cps/api/v1/payment/resp"

	service "github.com/mashbens/cps/business/payment"
	jwtService "github.com/mashbens/cps/business/user"
)

type PaymentController struct {
	paymentService service.PaymentService
	jwtService     jwtService.JWTService
}

func NewPaymentController(
	paymentService service.PaymentService,
	jwtService jwtService.JWTService,
) *PaymentController {
	return &PaymentController{
		paymentService: paymentService,
		jwtService:     jwtService,
	}
}

func (controller *PaymentController) CreatePayment(c echo.Context) error {
	var createPaymenReq request.CreatePaymentRequest
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
	if err := c.Bind(&createPaymenReq); err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	id := fmt.Sprintf("%v", claims["user_id"])
	intID, err := strconv.Atoi(id)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	createPaymenReq.UserID = intID

	payment, err := controller.paymentService.CreatePayment(request.NewCreatePaymentReq(createPaymenReq))
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	data := resp.FromService(*payment)
	response := _response.BuildSuccsessResponse("Payment Created", true, data)
	return c.JSON(http.StatusOK, response)
}

func (controller *PaymentController) GetPaymentDetail(c echo.Context) error {
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
	id := fmt.Sprintf("%v", claims["user_id"])

	payment, err := controller.paymentService.FindPaymentDetails(id)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	data := resp.FromService(*payment)

	response := _response.BuildSuccsessResponse("Payment found", true, data)
	return c.JSON(http.StatusOK, response)
}
