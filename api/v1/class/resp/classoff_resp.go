package resp

import (
	"github.com/mashbens/cps/business/class/entity"
)

type ClassResp struct {
	ID          int    `json:"id"`
	Classname   string `json:"classname"`
	Trainer     string `json:"trainer"`
	Date        string `json:"date"`
	Clock       string `json:"clock"`
	ClasType    string `json:"clastype"`
	Description string `json:"description"`
}

func FromService(class entity.Class) ClassResp {
	return ClassResp{
		ID:          class.ID,
		Classname:   class.Classname,
		Trainer:     class.Trainer,
		Date:        class.Date,
		Clock:       class.Clock,
		Description: class.Description,
		ClasType:    class.ClassType,
	}
}

func FromServiceSlice(data []entity.Class) []ClassResp {
	var memberAray []ClassResp
	for key := range data {
		memberAray = append(memberAray, FromService(data[key]))

	}
	return memberAray
}
