package entity

import (
	member "github.com/mashbens/cps/business/member/entity"
	user "github.com/mashbens/cps/business/user/entity"
)

type Payment struct {
	ID           int
	UserID       int
	User         user.User
	MembershipID int
	Membership   member.Membership
	Amount       int
	SnapURL      string
}
