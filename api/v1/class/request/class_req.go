package request

import (
	"mime/multipart"

	"github.com/mashbens/cps/business/class/entity"
)

type CreatClassReq struct {
	ID          int    `json:"id" form:"ID" `
	Classname   string `json:"classname" form:"classname"`
	Trainer     string `json:"trainer" form:"trainer"`
	Date        string `json:"date" form:"date"`
	Clock       string `json:"clock" form:"clock"`
	ClasType    string `json:"clastype" form:"clastype"`
	Capacity    int    `json:"capacity" form:"capacity"`
	Duration    int    `json:"duration" form:"duration"`
	Description string `json:"description" form:"description"`
	AdminID     int    `json:"admin_id" form:"admin_id"`
	Image       string `form:"image" `
	ImgBB       *multipart.FileHeader
}

func NewCreateClassReq(req CreatClassReq) entity.Class {
	return entity.Class{
		ID:          req.ID,
		Classname:   req.Classname,
		Trainer:     req.Trainer,
		Date:        req.Date,
		Clock:       req.Clock,
		Description: req.Description,
		Capacity:    req.Capacity,
		Duration:    req.Duration,
		ClassType:   req.ClasType,
		AdminID:     req.AdminID,
		ImgBB:       req.ImgBB,
	}
}
