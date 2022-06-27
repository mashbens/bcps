package resp

import (
	"github.com/mashbens/cps/business/admin/entity"
)

type AdminResp struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Token string `json:"token"`
}

func FromServiceAdmin(admin entity.Admin) AdminResp {
	return AdminResp{
		ID:    int(admin.ID),
		Name:  admin.Name,
		Email: admin.Email,
		Phone: admin.Phone,
		Token: admin.Token,
	}
}
