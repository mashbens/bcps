package member

import (
	"github.com/mashbens/cps/business/member"
	"github.com/mashbens/cps/business/member/entity"
	"gorm.io/gorm"
)

type MemberPostgresRepository struct {
	db *gorm.DB
}

func NewMemberPostgresRepository(db *gorm.DB) member.MemberRepo {
	return &MemberPostgresRepository{
		db: db,
	}
}

func (c *MemberPostgresRepository) FindMemberByID(memberID string) (entity.Membership, error) {
	var record Membership
	res := c.db.Where("id = ?", memberID).Take(&record)
	if res.Error != nil {
		return record.toService(), res.Error
	}
	return record.toService(), nil
}

func (c *MemberPostgresRepository) InserMemberships(member entity.Membership) (entity.Membership, error) {
	record := fromService(member)
	res := c.db.Create(&record)
	if res.Error != nil {
		return record.toService(), res.Error
	}
	return record.toService(), nil
}

func (c *MemberPostgresRepository) FIndAllMemberType(title string) (data []entity.Membership) {
	var record []Membership
	res := c.db.Find(&record)
	if res.Error != nil {
		return []entity.Membership{}
	}
	return toServiceList(record)
}
