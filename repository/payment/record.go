package payment

import (
	member "github.com/mashbens/cps/business/member/entity"
	"github.com/mashbens/cps/business/payment/entity"
	user "github.com/mashbens/cps/business/user/entity"
	// ----
)

type Payment struct {
	ID           int `gorm:"primary_key:auto_increment" json:"-"`
	UserID       int
	User         user.User `gorm:"foreignkey:UserID" json:"-"`
	MembershipID int
	Membership   member.Membership `gorm:"foreignkey:MembershipID" json:"-"`
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
