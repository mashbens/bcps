package request

import "github.com/mashbens/cps/business/superadmin/entity"

type SuperAdminReq struct {
	ID       int    `json:"id"`
	Name     string `json:"name" form:"name" binding:"required,min=1"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}

func NewSuperAdminReq(req SuperAdminReq) entity.SuperAdmin {
	return entity.SuperAdmin{
		ID:       req.ID,
		Name:     req.Name,
		Password: req.Password,
	}
}