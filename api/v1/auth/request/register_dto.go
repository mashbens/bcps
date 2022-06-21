package request

import (
	"github.com/mashbens/cps/business/user/entity"
)

type RegisterRequest struct {
	Name     string `json:"name" form:"name" binding:"required,min=1"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Phone    string `json:"phone" form:"phone" binding:"required,min=1"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}

func NewRegisterRequest(req RegisterRequest) entity.User {
	return entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: req.Password,
	}
}
