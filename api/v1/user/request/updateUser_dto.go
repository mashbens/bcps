package request

import "github.com/mashbens/cps/business/user/entity"

type UpdateUserRequest struct {
	ID       int    `json:"id"`
	Name     string `json:"name" form:"name" binding:"required,min=1"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Phone    string `json:"phone" form:"phone" binding:"required,min=1"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}

func NewUpdateUserRequest(req UpdateUserRequest) entity.User {
	return entity.User{
		ID:       req.ID,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Phone:    req.Phone,
	}
}
