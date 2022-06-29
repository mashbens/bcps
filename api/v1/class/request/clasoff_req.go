package request

import (
	"github.com/mashbens/cps/business/class/entity"
)

type CreatClassReq struct {
	ID          int    `json:"id"`
	Classname   string `json:"classname"`
	Trainer     string `json:"trainer"`
	Date        string `json:"date"`
	Clock       string `json:"clock"`
	ClasType    string `json:"clastype"`
	Description string `json:"description"`
	AdminID     int    `json:"admin_id"`
}

func NewCreateClassReq(req CreatClassReq) entity.Class {
	return entity.Class{
		ID:          req.ID,
		Classname:   req.Classname,
		Trainer:     req.Trainer,
		Date:        req.Date,
		Clock:       req.Clock,
		Description: req.Description,
		ClassType:   req.ClasType,
		AdminID:     req.AdminID,
	}
}
