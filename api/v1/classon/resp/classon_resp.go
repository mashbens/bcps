package resp

import (
	"github.com/mashbens/cps/business/classon/entity"
)

type ClassOnResp struct {
	ID          int    `json:"id"`
	Classname   string `json:"classname"`
	Trainer     string `json:"trainer"`
	Date        string `json:"date"`
	Clock       string `json:"clock"`
	Description string `json:"description"`
}

func FromService(class entity.Classon) ClassOnResp {
	return ClassOnResp{
		ID:          class.ID,
		Classname:   class.Classname,
		Trainer:     class.Trainer,
		Date:        class.Date,
		Clock:       class.Clock,
		Description: class.Description,
	}
}

func FromServiceSlice(data []entity.Classon) []ClassOnResp {
	var memberAray []ClassOnResp
	for key := range data {
		memberAray = append(memberAray, FromService(data[key]))

	}
	return memberAray
}
