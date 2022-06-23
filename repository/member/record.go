package member

import "github.com/mashbens/cps/business/member/entity"

type Membership struct {
	ID       int    `gorm:"primary_key:auto_increment" json:"-"`
	Type     string `gorm:"type:varchar(100)" `
	Price    int
	Duration int
}

func (m *Membership) toService() entity.Membership {
	return entity.Membership{
		ID:       int(m.ID),
		Type:     m.Type,
		Price:    m.Price,
		Duration: m.Duration,
	}
}

func fromService(member entity.Membership) Membership {
	return Membership{
		ID:       int(member.ID),
		Type:     member.Type,
		Price:    member.Price,
		Duration: member.Duration,
	}
}
