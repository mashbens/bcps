package request

import "github.com/mashbens/cps/business/member/entity"

type CreatMemberRequest struct {
	ID            int    `json:"ID"`
	Type          string `json:"type"`
	Price         int    `json:"price"`
	Duration      int    `json:"duration"`
	Super_adminID int    `json:"super_admin"`
}

func NewCreateMemberReq(req CreatMemberRequest) entity.Membership {
	return entity.Membership{
		ID:            req.ID,
		Type:          req.Type,
		Price:         req.Price,
		Duration:      req.Duration,
		Super_adminID: req.Super_adminID,
	}
}
