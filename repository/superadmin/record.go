package superadmin

import (
	"github.com/mashbens/cps/business/superadmin/entity"
)

type SuperAdmin struct {
	ID       int    `gorm:"primary_key:auto_increment" json:"-"`
	Name     string `gorm:"type:varchar(100)" json:"-"`
	Password string `gorm:"type:varchar(100)" json:"-"`
}

func (s *SuperAdmin) toService() entity.SuperAdmin {
	return entity.SuperAdmin{
		ID:       int(s.ID),
		Name:     s.Name,
		Password: s.Password,
	}
}

func fromService(s entity.SuperAdmin) SuperAdmin {
	return SuperAdmin{
		ID:       int(s.ID),
		Name:     s.Name,
		Password: s.Password,
	}
}
