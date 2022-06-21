package request

import "github.com/mashbens/cps/business/user/entity"

type PasswordResetRequest struct {
	Email string `json:"email" form:"email" binding:"required,email"`

	Password string `json:"password" form:"password" binding:"required,min=6"`
}

func NewPasswordResetRequest(req PasswordResetRequest) entity.User {
	return entity.User{
		Email:    req.Email,
		Password: req.Password,
	}
}
