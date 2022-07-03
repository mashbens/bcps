package request

import (
	"mime/multipart"

	"github.com/mashbens/cps/business/member/entity"
)

type CreatMemberRequest struct {
	ID            int    `form:"ID" json:"ID"`
	Type          string `form:"type" json:"type"`
	Price         int    `form:"price" json:"price"`
	Duration      int    `form:"duration" json:"duration"`
	Super_adminID int    `form:"super_admin" json:"super_admin"`
	Description   string `form:"description" json:"description"`
	Image         string `form:"image"`
	ImgBB         *multipart.FileHeader
}

func NewCreateMemberReq(req CreatMemberRequest) entity.Membership {
	return entity.Membership{
		ID:            req.ID,
		Type:          req.Type,
		Price:         req.Price,
		Duration:      req.Duration,
		Super_adminID: req.Super_adminID,
		Description:   req.Description,
		Img:           req.Image,
		ImgBB:         req.ImgBB,
	}
}
