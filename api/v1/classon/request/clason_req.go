package request

import (
	"github.com/mashbens/cps/business/classon/entity"
)

type CreatClassReq struct {
	ID          int    `json:"id"`
	Classname   string `json:"classname"`
	Trainer     string `json:"trainer"`
	Date        string `json:"date"`
	Clock       string `json:"clock"`
	Description string `json:"description"`
	AdminID     int    `json:"admin_id"`
}

func NewCreateClassReq(req CreatClassReq) entity.Classon {
	return entity.Classon{
		ID:          req.ID,
		Classname:   req.Classname,
		Trainer:     req.Trainer,
		Date:        req.Date,
		Clock:       req.Clock,
		Description: req.Description,
		AdminID:     req.AdminID,
	}
}