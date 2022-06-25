package resp

import "github.com/mashbens/cps/business/superadmin/entity"

type SuperAdminResp struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token,omitempty"`
}

func FromServiceSuperAdmin(s entity.SuperAdmin) SuperAdminResp {
	return SuperAdminResp{
		ID:    int(s.ID),
		Name:  s.Name,
		Token: s.Token,
	}
}
