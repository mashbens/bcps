package payment

import (
	"github.com/mashbens/cps/business/payment/entity"
)

type Payment struct {
	ID           int `gorm:"primary_key:auto_increment" json:"-"`
	UserID       int
	User         entity.User `gorm:"foreignkey:UserID" json:"-"`
	MembershipID int
	Membership   entity.Membership `gorm:"foreignkey:MembershipID" json:"-"`
	Amount       int
}

func (p *Payment) toService() entity.Payment {
	return entity.Payment{
		ID:           int(p.ID),
		UserID:       p.UserID,
		User:         p.User,
		MembershipID: p.MembershipID,
		Membership:   p.Membership,
		Amount:       p.Amount,
	}
}

func fromService(payment entity.Payment) Payment {
	return Payment{
		ID:           int(payment.ID),
		UserID:       payment.UserID,
		User:         payment.User,
		MembershipID: payment.MembershipID,
		Membership:   payment.Membership,
		Amount:       payment.Amount,
	}
}
