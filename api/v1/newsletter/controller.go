package newsletter

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"

	_response "github.com/mashbens/cps/api/common/response"
	"github.com/mashbens/cps/api/v1/newsletter/request"
	"github.com/mashbens/cps/api/v1/newsletter/resp"
	service "github.com/mashbens/cps/business/newsletter"
	jwtService "github.com/mashbens/cps/business/user"

	"github.com/labstack/echo/v4"
)

type NewsController struct {
	newsService service.NewsService
	jwtService  jwtService.JWTService
}

func NewNewsController(
	newsService service.NewsService,
	jwtService jwtService.JWTService,
) *NewsController {
	return &NewsController{
		newsService: newsService,
		jwtService:  jwtService,
	}
}

func (controller *NewsController) GetAllNews(c echo.Context) error {
	res := controller.newsService.FindAllNews("")

	data := resp.FromServiceSlice(res)

	_response := _response.BuildSuccsessResponse("All Newsletter found", true, data)
	return c.JSON(http.StatusOK, _response)
}

func (controller *NewsController) GetNewsByID(c echo.Context) error {

	id := c.Param("id")

	news, err := controller.newsService.FindNewsByID(id)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	data := resp.FromService(*news)
	_response := _response.BuildSuccsessResponse("Member found", true, data)
	return c.JSON(http.StatusOK, _response)
}

func (controller *NewsController) CreateNews(c echo.Context) error {
	var newNews request.CreatNewsReq
	header := c.Request().Header.Get("Authorization")
	err := c.Bind(&newNews)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if newNews.Title == "" || newNews.Content == "" || newNews.Date == "" {
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
		response := _response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	newNews.ImgBB = file
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	adminintID, err := strconv.Atoi(id)
	newNews.AdminID = adminintID

	News, err := controller.newsService.InsertNews(request.NewCreateNewsReq(newNews))
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	data := resp.FromService(*News)
	_response := _response.BuildSuccsessResponse("Newsletter created successfully", true, data)
	return c.JSON(http.StatusOK, _response)
}

func (controller *NewsController) UpdateNews(c echo.Context) error {
	var newNews request.CreatNewsReq
	header := c.Request().Header.Get("Authorization")
	err := c.Bind(&newNews)
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if newNews.Title == "" || newNews.Content == "" || newNews.Date == "" {
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
		response := _response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	newNews.ImgBB = file
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	adminintID, err := strconv.Atoi(id)
	newNews.AdminID = adminintID

	News, err := controller.newsService.UpdateNews(request.NewCreateNewsReq(newNews))
	if err != nil {
		response := _response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	data := resp.FromService(*News)
	_response := _response.BuildSuccsessResponse("Newsletter created successfully", true, data)
	return c.JSON(http.StatusOK, _response)
}

func (controller *NewsController) DeleteNews(c echo.Context) error {
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

	member := controller.newsService.DeleteNews(adminID, memberID)
	if member != nil {
		response := _response.BuildErrorResponse("Failed to process request", "Newsletter not found", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	_response := _response.BuildSuccsessResponse("Newsletter Deleted", true, nil)
	return c.JSON(http.StatusOK, _response)
}
