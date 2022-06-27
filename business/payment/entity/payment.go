package entity

import (
	_member "github.com/mashbens/cps/business/member/entity"
	_user "github.com/mashbens/cps/business/user/entity"
)

type Payment struct {
	ID           int
	UserID       int
	User         User
	MembershipID int
	Membership   Membership
	Amount       int
	SnapURL      string
}

type User struct {
	ID             int
	Name           string
	Email          string
	Phone          string
	Password       string
	Token          string
	Member_expired string
	Member_type    string
	Totp           string
}

type Membership struct {
	ID            int
	Type          string
	Price         int
	Duration      int
	Super_adminID int
}

func MemberToPayment(data _member.Membership) Membership {
	return Membership{
		ID:       data.ID,
		Type:     data.Type,
		Price:    data.Price,
		Duration: data.Duration,
	}
}

func UserToPayment(data _user.User) User {
	return User{
		ID:             data.ID,
		Name:           data.Name,
		Email:          data.Email,
		Phone:          data.Phone,
		Member_expired: data.Member_expired,
		Member_type:    data.Member_type,
	}
}
