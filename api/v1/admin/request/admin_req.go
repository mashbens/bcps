package request

import "github.com/mashbens/cps/business/admin/entity"

type AdminReq struct {
	ID           int    `json:"id"`
	Name         string `json:"name" form:"name" binding:"required,min=1"`
	Email        string `json:"email" form:"email" binding:"required,email"`
	Phone        string `json:"phone" form:"phone" binding:"required,min=1"`
	Password     string `json:"password" form:"password" binding:"required,min=6"`
	SuperAdminID int    `json:"superAdmin_id"`
}

func NewAdminReq(req AdminReq) entity.Admin {
	return entity.Admin{
		ID:           req.ID,
		Name:         req.Name,
		Email:        req.Email,
		Password:     req.Password,
		Phone:        req.Phone,
		SuperAdminID: req.SuperAdminID,
	}
}
