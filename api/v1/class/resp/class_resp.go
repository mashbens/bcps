package resp

import (
	"github.com/mashbens/cps/business/class/entity"
)

type ClassResp struct {
	ID          int    `json:"id"`
	Classname   string `json:"classname"`
	Status      string `json:"status"`
	Capacity    int    `json:"capacity"`
	Trainer     string `json:"trainer"`
	Date        string `json:"date"`
	Clock       string `json:"clock"`
	ClasType    string `json:"clastype"`
	Duration    int    `json:"duration"`
	UserBooked  int    `json:"user_booked"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func FromService(class entity.Class) ClassResp {
	return ClassResp{
		ID:          class.ID,
		Classname:   class.Classname,
		ClasType:    class.ClassType,
		Status:      class.Status,
		Capacity:    class.Capacity,
		Trainer:     class.Trainer,
		Date:        class.Date,
		Clock:       class.Clock,
		Duration:    class.Duration,
		Description: class.Description,
		UserBooked:  int(class.UserBooked),
		Image:       class.Img,
	}
}

func FromServiceSlice(data []entity.Class) []ClassResp {
	var memberAray []ClassResp
	for key := range data {
		memberAray = append(memberAray, FromService(data[key]))

	}
	return memberAray
}
